package heap

import "jvm/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (self *ClassMember) copyMemberInfo(info *classfile.MemberInfo) {
	self.accessFlags = info.AccessFlags()
	self.name = info.Name()
	self.descriptor = info.Descriptor()
}

func (self *ClassMember) IsPublic() bool       { return self.accessFlags&AccPublic != 0 }
func (self *ClassMember) IsPrivate() bool      { return self.accessFlags&AccPrivate != 0 }
func (self *ClassMember) IsProtected() bool    { return self.accessFlags&AccProtected != 0 }
func (self *ClassMember) IsStatic() bool       { return self.accessFlags&AccStatic != 0 }
func (self *ClassMember) IsFinal() bool        { return self.accessFlags&AccFinal != 0 }
func (self *ClassMember) IsSuper() bool        { return self.accessFlags&AccSuper != 0 }
func (self *ClassMember) IsSynchronized() bool { return self.accessFlags&AccSynchronized != 0 }
func (self *ClassMember) IsOpen() bool         { return self.accessFlags&AccOpen != 0 }
func (self *ClassMember) IsTransitive() bool   { return self.accessFlags&AccTransitive != 0 }
func (self *ClassMember) IsVolatile() bool     { return self.accessFlags&AccVolatile != 0 }
func (self *ClassMember) IsBridge() bool       { return self.accessFlags&AccBridge != 0 }
func (self *ClassMember) IsStaticPhase() bool  { return self.accessFlags&AccStaticPhase != 0 }
func (self *ClassMember) IsTransient() bool    { return self.accessFlags&AccTransient != 0 }
func (self *ClassMember) IsVarargs() bool      { return self.accessFlags&AccVarargs != 0 }
func (self *ClassMember) IsNative() bool       { return self.accessFlags&AccNative != 0 }
func (self *ClassMember) IsInterface() bool    { return self.accessFlags&AccInterface != 0 }
func (self *ClassMember) IsAbstract() bool     { return self.accessFlags&AccAbstract != 0 }
func (self *ClassMember) IsStrict() bool       { return self.accessFlags&AccStrict != 0 }
func (self *ClassMember) IsSynthetic() bool    { return self.accessFlags&AccSynthetic != 0 }
func (self *ClassMember) IsAnnotation() bool   { return self.accessFlags&AccAnnotation != 0 }
func (self *ClassMember) IsEnum() bool         { return self.accessFlags&AccEnum != 0 }
func (self *ClassMember) IsModule() bool       { return self.accessFlags&AccModule != 0 }
func (self *ClassMember) IsMandated() bool     { return self.accessFlags&AccMandated != 0 }

func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		return true
	}

	c := self.class
	if self.IsProtected() {
		return d == c || d.IsSubClassOf(c) || c.GetPackageName() == d.GetPackageName()
	}

	if !self.IsPrivate() {
		return c.GetPackageName() == d.GetPackageName()
	}

	return c == d
}
