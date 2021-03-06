# Gorm ForeignKey
Gorm don't manage the foreign key by default.

This package can. Basically, but it can do it.

Some things needs to be done to be perfectly compatible with gorm ; like using an another column name for the foreign key etc.

# Installation
```bash
go get -u github.com/youkoulayley/gormfk
```

# How to use it
You can define a new foreign key for this kind of relation : 

* Belongs To ;
* Many to Many.

Some will be coming in the future.

## Belongs To
After the automigrate functions, you can do this :

```go
import (
    'github.com/youkoulayley/gormfk'
)

...

Db().AutoMigrate(&models.Role{})
Db().AutoMigrate(&models.User{})

gormfk.BelongsTo(bootstrap.Db(), &models.User{}, &models.Role{})
```

## Many To Many
After the automigrate functions, you can do this :

```go
import (
    'github.com/youkoulayley/gormfk'
)

...

Db().AutoMigrate(&models.User{})
Db().AutoMigrate(&models.Channel{})

gormfk.Many2Many(bootstrap.Db(), &models.Channel{}, &models.User{})
```