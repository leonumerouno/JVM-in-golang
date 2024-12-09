package heap

import "jvm/classfile"

type Field struct {
	ClassMember
	constValueIndex uint
	slotId          uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}

func (self *Field) IsPublic() bool       { return self.accessFlags&AccPublic != 0 }
func (self *Field) IsPrivate() bool      { return self.accessFlags&AccPrivate != 0 }
func (self *Field) IsProtected() bool    { return self.accessFlags&AccProtected != 0 }
func (self *Field) IsStatic() bool       { return self.accessFlags&AccStatic != 0 }
func (self *Field) IsFinal() bool        { return self.accessFlags&AccFinal != 0 }
func (self *Field) IsSuper() bool        { return self.accessFlags&AccSuper != 0 }
func (self *Field) IsSynchronized() bool { return self.accessFlags&AccSynchronized != 0 }
func (self *Field) IsOpen() bool         { return self.accessFlags&AccOpen != 0 }
func (self *Field) IsTransitive() bool   { return self.accessFlags&AccTransitive != 0 }
func (self *Field) IsVolatile() bool     { return self.accessFlags&AccVolatile != 0 }
func (self *Field) IsBridge() bool       { return self.accessFlags&AccBridge != 0 }
func (self *Field) IsStaticPhase() bool  { return self.accessFlags&AccStaticPhase != 0 }
func (self *Field) IsTransient() bool    { return self.accessFlags&AccTransient != 0 }
func (self *Field) IsVarargs() bool      { return self.accessFlags&AccVarargs != 0 }
func (self *Field) IsNative() bool       { return self.accessFlags&AccNative != 0 }
func (self *Field) IsInterface() bool    { return self.accessFlags&AccInterface != 0 }
func (self *Field) IsAbstract() bool     { return self.accessFlags&AccAbstract != 0 }
func (self *Field) IsStrict() bool       { return self.accessFlags&AccStrict != 0 }
func (self *Field) IsSynthetic() bool    { return self.accessFlags&AccSynthetic != 0 }
func (self *Field) IsAnnotation() bool   { return self.accessFlags&AccAnnotation != 0 }
func (self *Field) IsEnum() bool         { return self.accessFlags&AccEnum != 0 }
func (self *Field) IsModule() bool       { return self.accessFlags&AccModule != 0 }
func (self *Field) IsMandated() bool     { return self.accessFlags&AccMandated != 0 }

func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (self *Field) ConstValueIndex() uint {
	return self.constValueIndex
}

func (self *Field) Descriptor() string {
	return self.descriptor
}

func (self *Field) Class() *Class {
	return self.class
}

func (self *Field) SlotId() uint {
	return self.slotId
}

func (self *Field) Name() string {
	return self.name
}
