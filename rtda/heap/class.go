package heap

import (
	"jvm/classfile"
	"strings"
)

type Class struct {
	accessFlags       uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
	initStarted       bool
	jClass            *Object
	sourceFile        string
}

const (
	jlObjectClassName       = "java/lang/Object"
	jlClassClassName        = "java/lang/Class"
	jlStringClassName       = "java/lang/String"
	jlThreadClassName       = "java/lang/Thread"
	jlCloneableClassName    = "java/lang/Cloneable"
	ioSerializableClassName = "java/io/Serializable"
)

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = NewConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = NewMethods(class, cf.Methods())
	class.sourceFile = getSourceFile(cf)
	return class
}

func (self *Class) IsPublic() bool       { return self.accessFlags&AccPublic != 0 }
func (self *Class) IsPrivate() bool      { return self.accessFlags&AccPrivate != 0 }
func (self *Class) IsProtected() bool    { return self.accessFlags&AccProtected != 0 }
func (self *Class) IsStatic() bool       { return self.accessFlags&AccStatic != 0 }
func (self *Class) IsFinal() bool        { return self.accessFlags&AccFinal != 0 }
func (self *Class) IsSuper() bool        { return self.accessFlags&AccSuper != 0 }
func (self *Class) IsSynchronized() bool { return self.accessFlags&AccSynchronized != 0 }
func (self *Class) IsOpen() bool         { return self.accessFlags&AccOpen != 0 }
func (self *Class) IsTransitive() bool   { return self.accessFlags&AccTransitive != 0 }
func (self *Class) IsVolatile() bool     { return self.accessFlags&AccVolatile != 0 }
func (self *Class) IsBridge() bool       { return self.accessFlags&AccBridge != 0 }
func (self *Class) IsStaticPhase() bool  { return self.accessFlags&AccStaticPhase != 0 }
func (self *Class) IsTransient() bool    { return self.accessFlags&AccTransient != 0 }
func (self *Class) IsVarargs() bool      { return self.accessFlags&AccVarargs != 0 }
func (self *Class) IsNative() bool       { return self.accessFlags&AccNative != 0 }
func (self *Class) IsInterface() bool    { return self.accessFlags&AccInterface != 0 }
func (self *Class) IsAbstract() bool     { return self.accessFlags&AccAbstract != 0 }
func (self *Class) IsStrict() bool       { return self.accessFlags&AccStrict != 0 }
func (self *Class) IsSynthetic() bool    { return self.accessFlags&AccSynthetic != 0 }
func (self *Class) IsAnnotation() bool   { return self.accessFlags&AccAnnotation != 0 }
func (self *Class) IsEnum() bool         { return self.accessFlags&AccEnum != 0 }
func (self *Class) IsModule() bool       { return self.accessFlags&AccModule != 0 }
func (self *Class) IsMandated() bool     { return self.accessFlags&AccMandated != 0 }

func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() || self.GetPackageName() == other.GetPackageName()
}

func (self *Class) GetPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) IsSubClassOf(c *Class) bool {
	for k := self.superClass; k != nil; k = k.superClass {
		if k == c {
			return true
		}
	}
	return false
}

func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}

func (self *Class) StaticVars() Slots {
	return self.staticVars
}

func (self *Class) isAssignableFrom(other *Class) bool {
	s, t := other, self
	if s == t {
		return true
	}
	if !s.IsArray() {
		if !s.IsInterface() {
			if !t.IsInterface() {
				return s.IsSubClassOf(t)
			} else {
				return s.IsImplements(t)
			}
		} else {
			if !t.IsInterface() {
				return t.IsJlObject()
			} else {
				return s.isSubInterfaceOf(t)
			}
		}
	} else {
		if !t.IsArray() {
			if !t.IsInterface() {
				return t.IsJlObject()
			} else {
				return t.IsJlCloneable() || t.IsJioSerializable()
			}
		} else {
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.isAssignableFrom(sc)
		}
	}
}

func (self *Class) IsJlObject() bool {
	return self.name == jlObjectClassName
}

func (self *Class) IsJlCloneable() bool {
	return self.name == jlCloneableClassName
}

func (self *Class) IsJioSerializable() bool {
	return self.name == ioSerializableClassName
}

// iface extends class
func (self *Class) isSuperInterfaceOf(iface *Class) bool {
	return iface.isSubInterfaceOf(self)
}

func (self *Class) IsImplements(other *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == other || i.isSubInterfaceOf(other) {
				return true
			}
		}
	}
	return false
}

func (self *Class) isSubInterfaceOf(other *Class) bool {
	for _, superInterface := range other.interfaces {
		if superInterface == other || superInterface.isSubInterfaceOf(other) {
			return true
		}
	}
	return false
}

func (self *Class) GetMainMethod() *Method {
	return self.GetStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) GetStaticMethod(name, descriptor string) *Method {
	return self.getMethod(name, descriptor, true)
}

func (self *Class) GetInstanceMethod(name, descriptor string) *Method {
	return self.getMethod(name, descriptor, false)
}

func (class *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for k := class; k != nil; k = k.SuperClass() {
		for _, method := range k.Methods() {
			if method.IsStatic() == isStatic &&
				method.Name() == name &&
				method.Descriptor() == descriptor {
				return method
			}
		}
	}
	return nil
}

func (self *Class) SuperClass() *Class {
	return self.superClass
}

func (self *Class) Name() string {
	return self.name
}

func (self *Class) InitStarted() bool {
	return self.initStarted
}

func (self *Class) StartInit() {
	self.initStarted = true
}

func (self *Class) GetClinitMethod() *Method {
	return self.GetStaticMethod("<clinit>", "()V")
}

func (self *Class) Loader() *ClassLoader {
	return self.loader
}

func (self *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(self.name)
	return self.loader.LoadClass(arrayClassName)
}

func (self *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := self; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic && field.name == name && field.descriptor == descriptor {
				return field
			}
		}
	}
	return nil
}

func (self *Class) JClass() *Object {
	return self.jClass
}

func (self *Class) JavaName() string {
	return strings.Replace(self.name, "/", ".", -1)
}

func (self *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[self.name]
	return ok
}

func (self *Class) GetRefVar(name, descriptor string) *Object {
	field := self.getField(name, descriptor, true)
	slots := self.StaticVars()
	return slots.GetRef(field.slotId)
}

func (self *Class) Methods() []*Method {
	return self.methods
}

func (self *Class) SourceFile() string {
	return self.sourceFile
}

func getSourceFile(cf *classfile.ClassFile) string {
	if sfAttr := cf.SourceFileAttribute(); sfAttr != nil {
		return sfAttr.FileName()
	}
	return "Unknown"
}
