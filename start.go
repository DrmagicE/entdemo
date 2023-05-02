package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"entdemo/ent"
	"entdemo/ent/migrate"
	_ "entdemo/ent/runtime"
	"entdemo/ent/schema"
	"entdemo/ent/user"
)

func main() {
	client, err := ent.Open(
		"sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1",
		// 想看建表语句，打开下面的注释
		//ent.Debug(),
	)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// client.Schema.Create 会自动建表，示例中使用内存模式的数据库，所以每次都会重新建表。
	if err := client.Schema.Create(
		context.Background(),
		// ent默认会为关联字段创建外键，按使用习惯决定是否增加这行配置禁用外键。
		migrate.WithForeignKeys(false),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	//AddCarsToUser(client)
	//SetOwnerToCar(client)

	//AddUserToGroup(client)
	// AddGroupToUser(client)

	//QueryUserCars(client)
	//QueryUserWithCarsEagerLoading(client)

	// 软删除&硬删除
	//DeleteUser(client)
	// DeleteUserForceDelete(client)

	// 事物
	//GenTx(context.Background(), client)
	//WithTx(context.Background(), client.Debug(), func(tx *ent.Tx) error {
	//	CreateAndQueryUser(tx.Client())
	//	AddCarsToUser(tx.Client())
	//	return nil
	//})

}

func CreateAndQueryUser(client *ent.Client) {
	// 创建 name="user1"的用户
	client.Debug().User.Create().SetName("user1").SaveX(context.Background())
	// 查询 name="user1" 的用户
	fmt.Println(client.Debug().User.Query().Where(user.Name("user1")).All(context.Background()))
}

func AddCarsToUser(client *ent.Client) {
	// 创建2辆车
	car1 := client.Debug().Car.Create().SetName("car1").SaveX(context.Background())
	car2 := client.Debug().Car.Create().SetName("car2").SaveX(context.Background())
	// 这两辆车属于 name="user2"的用户
	client.Debug().User.Create().SetName("user2").AddCars(car1, car2).SaveX(context.Background())
}

func SetOwnerToCar(client *ent.Client) {
	// 创建一个用户
	user := client.Debug().User.Create().SetName("user2").SaveX(context.Background())
	// 创建车，并绑定用户
	client.Debug().Car.Create().SetName("car3").SetOwner(user).SaveX(context.Background())
}

func AddUserToGroup(client *ent.Client) {
	// 创建用户
	user1 := client.Debug().User.Create().SetName("user1").SaveX(context.Background())
	user2 := client.Debug().User.Create().SetName("user2").SaveX(context.Background())

	// 创建用户组，将用户加入到用户组
	client.Debug().Group.Create().SetName("group1").AddUsers(user1, user2).SaveX(context.Background())
}

func AddGroupToUser(client *ent.Client) {
	// 创建用户组
	group1 := client.Debug().Group.Create().SetName("group1").SaveX(context.Background())
	group2 := client.Debug().Group.Create().SetName("group2").SaveX(context.Background())

	// 用户加入用户组
	client.Debug().User.Create().SetName("user1").AddGroups(group1, group2).SaveX(context.Background())
}

func QueryUserCars(client *ent.Client) {

	// 初始化测试数据，创建用户和车
	u := client.User.Create().SetName("user1").SaveX(context.Background())
	client.Car.Create().SetName("car1").SetOwner(u).SaveX(context.Background())
	// 以下两种方法都可
	//car := client.Debug().User.QueryCars(u).AllX(context.Background())
	car := client.Debug().User.Query().Where(user.Name("user1")).QueryCars().AllX(context.Background())

	fmt.Println(car)
}

func QueryUserWithCarsEagerLoading(client *ent.Client) {
	// 初始化测试数据，创建用户和车
	u := client.User.Create().SetName("user1").SaveX(context.Background())
	client.Car.Create().SetName("car1").SetOwner(u).SaveX(context.Background())

	// 使用WithXXX()方法实现即时加载
	u = client.Debug().User.Query().WithCars().OnlyX(context.Background())
	fmt.Println(u.Edges.CarsOrErr())
}

func DeleteUser(client *ent.Client) {
	// 初始化测试数据
	client.User.Create().SetName("user1").SaveX(context.Background())

	client.Debug().User.Delete().Where(user.Name("user1")).ExecX(context.Background())
}

func DeleteUserForceDelete(client *ent.Client) {
	// 初始化测试数据
	client.User.Create().SetName("user1").SaveX(context.Background())

	ctx := context.Background()
	// 跳过软删除，直接删除数据库记录
	ctx = schema.SkipSoftDelete(ctx)
	client.Debug().User.Delete().Where(user.Name("user1")).ExecX(ctx)
}

func GenTx(ctx context.Context, client *ent.Client) error {
	client = client.Debug()
	// 开启事务
	tx, err := client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("starting a transaction: %w", err)
	}
	_, err = tx.User.Create().SetName("user1").Save(ctx)
	if err != nil {
		// 失败回滚
		return rollback(tx, err)
	}
	_, err = tx.Car.Create().SetName("car1").Save(ctx)
	if err != nil {
		return rollback(tx, err)
	}
	// 提交事务
	return tx.Commit()
}

// rollback 事务回滚方法
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
