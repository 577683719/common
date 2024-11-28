package condition_gen

import (
	"encoding/json"
	"fmt"
	set "github.com/duke-git/lancet/v2/datastructure/set"
	"github.com/shopspring/decimal"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"text/template"
	"time"
)

// 模板字符串
// 定义模板字符串

const tmpl = `
package {{.PackageName}}

import (
	{{- range .Imports }}
	"{{.}}"
	{{- end }}
	"time"
	"ecs_cloud/service_api/repository/dao"
	"github.com/shopspring/decimal"
)



{{- range .Methods }}
 {{ $method := . }}
func GetCondition(source {{.SourceOldType}}) dao.Condition {
	condition := dao.Condition{}
	{{- range .Fields }}
	condition.AndWithCondition(isNotDefautValue(source.{{.SourceField}}),{{.FieldClo}},"=",source.{{.SourceField}})
	{{- end }}

	return condition
}

{{- end }}

{{- range .Methods }}
 {{ $method := . }}
func GetPointerCondition(source *{{.SourceOldType}}) dao.Condition {
	condition := dao.Condition{}

	{{- range .Fields }}
	condition.AndWithCondition(isNotDefautValue(source.{{.SourceField}}),{{.FieldClo}},"=",source.{{.SourceField}})
	{{- end }}

	return condition
}

{{- end }}

func isNotDefautValue(value interface{}) bool {
	// 根据需要进行类型转换的逻辑
	switch v := value.(type) {
	case time.Time:
		return !v.IsZero() // 时间类型不为零则有效
	case string:
		return v != "" // 字符串非空则有效
	case int:
		return v != 0 // 整数不为零则有效
	case int8:
		return v != 0 // 整数不为零则有效
	case int16:
		return v != 0 // 整数不为零则有效
	case int32:
		return v != 0 // 整数不为零则有效
	case int64:
		return v != 0 // 整数不为零则有效
	case decimal.Decimal:
		return !v.IsZero()
	default:
		return false // 其他类型不去添加condition
	}
}
`

// 创建一个方法字符串转时间"stringToDecimal": StringToDecimal,
func stringToTime(str string) (time.Time, error) {
	return time.Parse(time.RFC3339, str)
}

//"yyyy-mm-dd hh:mm:ss"格式字符串转时间

func convertField(value interface{}) string {
	// 根据需要进行类型转换的逻辑
	switch v := value.(type) {
	case time.Time:
		return v.Format("2006-01-02 15:04:05")
	case decimal.Decimal:
		return v.String()
	}
	return fmt.Sprintf("%v", value)
}

type TemplateData struct {
	PathName           string
	PackageName        string
	Methods            []MethodData
	Imports            []string
	TypeToStringFileds []string
	IgnoreFields       []string
}

func NewTemplateData() TemplateData {
	return TemplateData{}
}
func (t *TemplateData) SetTypeToStringFileds(typeToString ...string) *TemplateData {
	t.TypeToStringFileds = typeToString
	return t
}
func (t *TemplateData) SetIgnoreFields(ignoreFields ...string) *TemplateData {
	t.IgnoreFields = ignoreFields
	return t
}
func (t *TemplateData) SetPackageName(packageName string) *TemplateData {
	t.PackageName = packageName
	return t
}
func (t *TemplateData) SetPathName(pathName string) *TemplateData {
	t.PathName = pathName
	return t
}
func (t *MethodData) SetIgnoreFields(ignoreFields ...string) *MethodData {
	t.IgnoreFields = ignoreFields
	return t
}
func (t *MethodData) SetTypeToStringFileds(typeToStringFileds ...string) *MethodData {
	t.TypeToStringFileds = typeToStringFileds
	return t
}

func (t *MethodData) SetStringToTimeFields(StringToTimeFields ...string) *MethodData {
	t.StringToTimeFields = StringToTimeFields
	return t
}
func (t *MethodData) SetStringToDecimalFileds(StringToDecimalFileds ...string) *MethodData {
	t.StringToDecimalFileds = StringToDecimalFileds
	return t
}

// 给TemplateData里面的字段生成set方法
// 给PackageName Methods Imports TypeToStringFileds IgnoreFields生成Set方法
type MethodData struct {
	ConvertName           string
	MethodName            string
	MethodListName        string
	SourceType            string
	SourceOldType         string
	TargetType            string
	Column                string
	TargetOldType         string
	Fields                []Field
	IgnoreFields          []string
	StringToTimeFields    []string
	TypeToStringFileds    []string
	StringToDecimalFileds []string
	RetruenName           string
}

type Field struct {
	SourceField string
	TargetField string
	Type        string
	FieldClo    string
}

// ShouldIgnore 方法用于判断某个字段是否在忽略列表中
func ShouldIgnore(s []string, field string) bool {

	for _, ignoreField := range s {
		if field == ignoreField {
			return true
		}
	}
	return false
}

// typeConverStr 方法用于判断某个字段是否在忽略列表中
func TypeConver(s []string, field string) bool {
	for _, typeConverField := range s {
		fmt.Println("typeConverField=" + typeConverField)
		fmt.Println("field=" + field)
		if field == typeConverField {
			return true
		}
	}
	return false
}

func StringToTime(s []string, field string) bool {
	for _, typeConverField := range s {
		if field == typeConverField {
			return true
		}
	}
	return false
}
func StringToDecimal(s []string, field string) bool {
	for _, stringToDecimal := range s {
		if field == stringToDecimal {
			return true
		}
	}
	return false
}
func generateFields(source interface{}, target interface{}) ([]Field, error) {
	sourceValue := reflect.ValueOf(source)
	targetValue := reflect.ValueOf(target)

	// Handle pointers
	if sourceValue.Kind() == reflect.Ptr {
		sourceValue = sourceValue.Elem()
	}
	if targetValue.Kind() == reflect.Ptr {
		targetValue = targetValue.Elem()
	}

	if sourceValue.Kind() != reflect.Struct || targetValue.Kind() != reflect.Struct {
		return nil, fmt.Errorf("both source and target must be structs")
	}

	sourceType := sourceValue.Type()

	targetType := targetValue.Type()
	fields := make([]Field, 0)

	sourceFieldMap := make(map[string]reflect.StructField)
	for i := 0; i < sourceType.NumField(); i++ {
		sourceField := sourceType.Field(i)
		sourceFieldMap[sourceField.Name] = sourceField
	}

	for i := 0; i < targetType.NumField(); i++ {
		targetField := targetType.Field(i)
		for sourceName, sourceField := range sourceFieldMap {
			if targetField.Name == sourceField.Name { // Match by type for simplicity
				field := Field{
					SourceField: sourceName,
					TargetField: targetField.Name,
					Type:        sourceField.Type.String(),
					FieldClo:    strings.Replace(sourceType.String(), "*", "", -1) + "Columns." + targetField.Name,
				}
				fields = append(fields, field)
				break
			}
		}
	}

	return fields, nil
}

// convertFilePathToPackagePath 将文件系统路径转换成包路径
func convertFilePathToPackagePath(filePath string) (string, error) {
	// 将反斜杠替换为正斜杠，以统一路径分隔符
	filePath = filepath.ToSlash(filePath)

	// 查找 ecs_cloud 的起始位置
	index := strings.Index(filePath, "ecs_cloud")
	if index == -1 {
		return "", fmt.Errorf("ecs_cloud not found in file path")
	}

	// 保留 ecs_cloud 及其之后的部分作为包路径
	packagePath := filePath[index:]

	return packagePath, nil
}

func getStructPackagePath(i interface{}) (string, error) {
	t := reflect.TypeOf(i)
	if t == nil {
		return "", fmt.Errorf("nil interface provided")
	}

	// 获取结构体类型的方法，以确保我们能找到它的包路径
	method := t.Method(0)
	pc := method.Func.Pointer()

	// 使用 runtime.FuncForPC 获取函数的文件路径
	fn := runtime.FuncForPC(pc)
	filePath, _ := fn.FileLine(pc)

	// 从文件路径中提取包路径
	dir := filepath.Dir(filePath)

	// 进一步处理以获得 Go 语言包路径，而不是文件系统路径
	goSrcIndex := strings.Index(dir, "src/")
	if goSrcIndex != -1 {
		dir = dir[goSrcIndex+4:] // 去除 "src/" 部分
	} else {
		// 如果在 Go modules 中，处理 /pkg/mod/ 中的路径
		modIndex := strings.Index(dir, "/pkg/mod/")
		if modIndex != -1 {
			dir = dir[modIndex+9:] // 去除 "/pkg/mod/" 部分
			// 去除版本信息，例如 /@v0.0.0-20210101000000-000000000000
			if atIndex := strings.LastIndex(dir, "@"); atIndex != -1 {
				dir = dir[:atIndex]
			}
		}
	}

	return dir, nil
}

func getPackage(i interface{}) string {
	path, _ := getStructPackagePath(i)
	packagePath, _ := convertFilePathToPackagePath(path)
	return packagePath
}

func BuildMethodData(sources []interface{}, target []interface{}) []MethodData {
	if len(sources) != len(target) {
		panic("缺少参数")
	}
	methodData := []MethodData{}
	for i, _ := range sources {

		sourceType := reflect.TypeOf(sources[i]).String()
		targetType := reflect.TypeOf(target[i]).String()
		retruenName := targetType[1:]
		sourceIndex := strings.LastIndex(sourceType, ".")
		targetIndex := strings.LastIndex(targetType, ".")
		sourceName := sourceType[sourceIndex+1:]
		targetName := targetType[targetIndex+1:]
		fields, _ := generateFields(sources[i], target[i])
		methodData = append(methodData,
			MethodData{
				MethodName:         fmt.Sprintf("%vTo%vConvert", sourceName, targetName),
				MethodListName:     fmt.Sprintf("%vListTo%vListConvert", sourceName, targetName),
				SourceType:         sourceType,
				SourceOldType:      strings.Replace(sourceType, "*", "", -1),
				TargetType:         targetType,
				TargetOldType:      strings.Replace(targetType, "*", "", -1),
				Fields:             fields,
				IgnoreFields:       nil,
				TypeToStringFileds: nil,
				RetruenName:        retruenName,
				//{{.SourceOldType}}Columns.{{.SourceField}}
				Column: strings.Replace(sourceType, "*", "", -1) + "Columns.",
			},
		)
	}
	return methodData
}
func Buildimports(sources []interface{}, target []interface{}) []string {
	if len(sources) != len(target) {
		panic("缺少参数")
	}
	imports := []string{}
	for i, _ := range sources {
		imports = append(imports,
			getPackage(sources[i]),
			getPackage(target[i]),
		)
	}
	return set.FromSlice(imports).ToSlice()
}
func BuildTemplateData(PackageName string, sources []interface{}, target []interface{}) TemplateData {
	fmt.Println("IgnoreFields")
	imports := Buildimports(sources, target)
	return TemplateData{
		PackageName: PackageName,
		Imports:     imports,
	}
}
func (t *TemplateData) SetMethodData(methods []MethodData) {
	t.Methods = methods
}
func (t *TemplateData) GenConevert() {
	tmpl, err := template.New("convert").Funcs(template.FuncMap{
		"shouldIgnore":    ShouldIgnore,
		"typeConver":      TypeConver,
		"stringToTime":    StringToTime,
		"stringToDecimal": StringToDecimal,
	}).Parse(tmpl)

	if err != nil {
		panic(err)
	}

	// 获取目录路径
	dirPath := filepath.Dir(t.PathName)

	// 检查目录是否存在，不存在则创建
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		panic(fmt.Sprintf("Failed to create directory: %s", err))
	}
	//f, err := os.Create("service/convert/orders_convert.go")
	f, err := os.Create(t.PathName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, t)
	if err != nil {
		panic(err)
	}
}
func (t *TemplateData) genConevert(pathName, packageName string, sourse, target []interface{}) {
	fmt.Println(t.IgnoreFields)
	data := BuildTemplateData(packageName, sourse, target)
	fmt.Println(data.TypeToStringFileds)
	marshal, _ := json.Marshal(data)
	fmt.Println("-------------------")
	fmt.Println(string(marshal))
	fmt.Println("-------------------")
	tmpl, err := template.New("convert").Funcs(template.FuncMap{
		"shouldIgnore":    ShouldIgnore,
		"typeConver":      TypeConver,
		"stringToTime":    StringToTime,
		"stringToDecimal": StringToDecimal,
	}).Parse(tmpl)

	if err != nil {
		panic(err)
	}
	//f, err := os.Create("service/convert/orders_convert.go")
	f, err := os.Create(pathName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		panic(err)
	}
}
