package gormfk

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/inflection"
)

// BelongsTo create the foreign key for a belongs to relation
func BelongsTo(db *gorm.DB, parentModel interface{}, childModel interface{}) {
	// Get name of the model
	tableParentAccessor := ReduceModelToName(parentModel)
	tableChildAccessor := ReduceModelToName(childModel)

	// Set model name to plural
	tableParentName := inflection.Plural(tableParentAccessor)
	tableChildName := inflection.Plural(tableChildAccessor)

	db.Table(tableParentName).AddForeignKey(tableChildAccessor+"_id", tableChildName+"(id)", "CASCADE", "CASCADE")
}

// Many2Many create the index for many to many relations
func Many2Many(db *gorm.DB, parentModel interface{}, childModel interface{}) {
	// Get name of the model
	tableParentAccessor := ReduceModelToName(parentModel)
	tableChildAccessor := ReduceModelToName(childModel)

	// Set model name to plural
	tableParentName := inflection.Plural(tableParentAccessor)
	tableChildName := inflection.Plural(tableChildAccessor)

	// Join the tables
	joinTable := fmt.Sprintf("%s_%s", tableParentAccessor, tableChildAccessor)

	// Add foreign Key and unique index
	db.Table(joinTable).AddForeignKey(tableParentAccessor+"_id", tableParentName+"(id)", "CASCADE", "CASCADE")
	db.Table(joinTable).AddForeignKey(tableChildAccessor+"_id", tableChildName+"(id)", "CASCADE", "CASCADE")
	db.Table(joinTable).AddUniqueIndex(joinTable+"_unique", tableParentAccessor+"_id", tableChildAccessor+"_id")
}

// ReduceModelToName get the name of the model by its reference
func ReduceModelToName(model interface{}) string {
	value := reflect.ValueOf(model)
	if value.Kind() != reflect.Ptr {
		return ""
	}

	elem := value.Elem()
	t := elem.Type()
	rawName := t.Name()
	return gorm.ToDBName(rawName)
}
