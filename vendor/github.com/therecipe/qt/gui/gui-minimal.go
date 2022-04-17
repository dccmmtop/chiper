// +build minimal

package gui

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "gui-minimal.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtGui_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtGui_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return []byte(gs)
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

//go:generate stringer -type=QAccessible__Event
//QAccessible::Event
type QAccessible__Event int64

var (
	QAccessible__SoundPlayed                     QAccessible__Event = QAccessible__Event(0x0001)
	QAccessible__Alert                           QAccessible__Event = QAccessible__Event(0x0002)
	QAccessible__ForegroundChanged               QAccessible__Event = QAccessible__Event(0x0003)
	QAccessible__MenuStart                       QAccessible__Event = QAccessible__Event(0x0004)
	QAccessible__MenuEnd                         QAccessible__Event = QAccessible__Event(0x0005)
	QAccessible__PopupMenuStart                  QAccessible__Event = QAccessible__Event(0x0006)
	QAccessible__PopupMenuEnd                    QAccessible__Event = QAccessible__Event(0x0007)
	QAccessible__ContextHelpStart                QAccessible__Event = QAccessible__Event(0x000C)
	QAccessible__ContextHelpEnd                  QAccessible__Event = QAccessible__Event(0x000D)
	QAccessible__DragDropStart                   QAccessible__Event = QAccessible__Event(0x000E)
	QAccessible__DragDropEnd                     QAccessible__Event = QAccessible__Event(0x000F)
	QAccessible__DialogStart                     QAccessible__Event = QAccessible__Event(0x0010)
	QAccessible__DialogEnd                       QAccessible__Event = QAccessible__Event(0x0011)
	QAccessible__ScrollingStart                  QAccessible__Event = QAccessible__Event(0x0012)
	QAccessible__ScrollingEnd                    QAccessible__Event = QAccessible__Event(0x0013)
	QAccessible__MenuCommand                     QAccessible__Event = QAccessible__Event(0x0018)
	QAccessible__ActionChanged                   QAccessible__Event = QAccessible__Event(0x0101)
	QAccessible__ActiveDescendantChanged         QAccessible__Event = QAccessible__Event(0x0102)
	QAccessible__AttributeChanged                QAccessible__Event = QAccessible__Event(0x0103)
	QAccessible__DocumentContentChanged          QAccessible__Event = QAccessible__Event(0x0104)
	QAccessible__DocumentLoadComplete            QAccessible__Event = QAccessible__Event(0x0105)
	QAccessible__DocumentLoadStopped             QAccessible__Event = QAccessible__Event(0x0106)
	QAccessible__DocumentReload                  QAccessible__Event = QAccessible__Event(0x0107)
	QAccessible__HyperlinkEndIndexChanged        QAccessible__Event = QAccessible__Event(0x0108)
	QAccessible__HyperlinkNumberOfAnchorsChanged QAccessible__Event = QAccessible__Event(0x0109)
	QAccessible__HyperlinkSelectedLinkChanged    QAccessible__Event = QAccessible__Event(0x010A)
	QAccessible__HypertextLinkActivated          QAccessible__Event = QAccessible__Event(0x010B)
	QAccessible__HypertextLinkSelected           QAccessible__Event = QAccessible__Event(0x010C)
	QAccessible__HyperlinkStartIndexChanged      QAccessible__Event = QAccessible__Event(0x010D)
	QAccessible__HypertextChanged                QAccessible__Event = QAccessible__Event(0x010E)
	QAccessible__HypertextNLinksChanged          QAccessible__Event = QAccessible__Event(0x010F)
	QAccessible__ObjectAttributeChanged          QAccessible__Event = QAccessible__Event(0x0110)
	QAccessible__PageChanged                     QAccessible__Event = QAccessible__Event(0x0111)
	QAccessible__SectionChanged                  QAccessible__Event = QAccessible__Event(0x0112)
	QAccessible__TableCaptionChanged             QAccessible__Event = QAccessible__Event(0x0113)
	QAccessible__TableColumnDescriptionChanged   QAccessible__Event = QAccessible__Event(0x0114)
	QAccessible__TableColumnHeaderChanged        QAccessible__Event = QAccessible__Event(0x0115)
	QAccessible__TableModelChanged               QAccessible__Event = QAccessible__Event(0x0116)
	QAccessible__TableRowDescriptionChanged      QAccessible__Event = QAccessible__Event(0x0117)
	QAccessible__TableRowHeaderChanged           QAccessible__Event = QAccessible__Event(0x0118)
	QAccessible__TableSummaryChanged             QAccessible__Event = QAccessible__Event(0x0119)
	QAccessible__TextAttributeChanged            QAccessible__Event = QAccessible__Event(0x011A)
	QAccessible__TextCaretMoved                  QAccessible__Event = QAccessible__Event(0x011B)
	QAccessible__TextColumnChanged               QAccessible__Event = QAccessible__Event(0x011D)
	QAccessible__TextInserted                    QAccessible__Event = QAccessible__Event(0x011E)
	QAccessible__TextRemoved                     QAccessible__Event = QAccessible__Event(0x011F)
	QAccessible__TextUpdated                     QAccessible__Event = QAccessible__Event(0x0120)
	QAccessible__TextSelectionChanged            QAccessible__Event = QAccessible__Event(0x0121)
	QAccessible__VisibleDataChanged              QAccessible__Event = QAccessible__Event(0x0122)
	QAccessible__ObjectCreated                   QAccessible__Event = QAccessible__Event(0x8000)
	QAccessible__ObjectDestroyed                 QAccessible__Event = QAccessible__Event(0x8001)
	QAccessible__ObjectShow                      QAccessible__Event = QAccessible__Event(0x8002)
	QAccessible__ObjectHide                      QAccessible__Event = QAccessible__Event(0x8003)
	QAccessible__ObjectReorder                   QAccessible__Event = QAccessible__Event(0x8004)
	QAccessible__Focus                           QAccessible__Event = QAccessible__Event(0x8005)
	QAccessible__Selection                       QAccessible__Event = QAccessible__Event(0x8006)
	QAccessible__SelectionAdd                    QAccessible__Event = QAccessible__Event(0x8007)
	QAccessible__SelectionRemove                 QAccessible__Event = QAccessible__Event(0x8008)
	QAccessible__SelectionWithin                 QAccessible__Event = QAccessible__Event(0x8009)
	QAccessible__StateChanged                    QAccessible__Event = QAccessible__Event(0x800A)
	QAccessible__LocationChanged                 QAccessible__Event = QAccessible__Event(0x800B)
	QAccessible__NameChanged                     QAccessible__Event = QAccessible__Event(0x800C)
	QAccessible__DescriptionChanged              QAccessible__Event = QAccessible__Event(0x800D)
	QAccessible__ValueChanged                    QAccessible__Event = QAccessible__Event(0x800E)
	QAccessible__ParentChanged                   QAccessible__Event = QAccessible__Event(0x800F)
	QAccessible__HelpChanged                     QAccessible__Event = QAccessible__Event(0x80A0)
	QAccessible__DefaultActionChanged            QAccessible__Event = QAccessible__Event(0x80B0)
	QAccessible__AcceleratorChanged              QAccessible__Event = QAccessible__Event(0x80C0)
	QAccessible__InvalidEvent                    QAccessible__Event = QAccessible__Event(C.QAccessible_InvalidEvent_Type())
)

//go:generate stringer -type=QAccessible__InterfaceType
//QAccessible::InterfaceType
type QAccessible__InterfaceType int64

const (
	QAccessible__TextInterface         QAccessible__InterfaceType = QAccessible__InterfaceType(0)
	QAccessible__EditableTextInterface QAccessible__InterfaceType = QAccessible__InterfaceType(1)
	QAccessible__ValueInterface        QAccessible__InterfaceType = QAccessible__InterfaceType(2)
	QAccessible__ActionInterface       QAccessible__InterfaceType = QAccessible__InterfaceType(3)
	QAccessible__ImageInterface        QAccessible__InterfaceType = QAccessible__InterfaceType(4)
	QAccessible__TableInterface        QAccessible__InterfaceType = QAccessible__InterfaceType(5)
	QAccessible__TableCellInterface    QAccessible__InterfaceType = QAccessible__InterfaceType(6)
)

//go:generate stringer -type=QAccessible__RelationFlag
//QAccessible::RelationFlag
type QAccessible__RelationFlag int64

const (
	QAccessible__Label        QAccessible__RelationFlag = QAccessible__RelationFlag(0x00000001)
	QAccessible__Labelled     QAccessible__RelationFlag = QAccessible__RelationFlag(0x00000002)
	QAccessible__Controller   QAccessible__RelationFlag = QAccessible__RelationFlag(0x00000004)
	QAccessible__Controlled   QAccessible__RelationFlag = QAccessible__RelationFlag(0x00000008)
	QAccessible__AllRelations QAccessible__RelationFlag = QAccessible__RelationFlag(0xffffffff)
)

//go:generate stringer -type=QAccessible__Role
//QAccessible::Role
type QAccessible__Role int64

const (
	QAccessible__NoRole               QAccessible__Role = QAccessible__Role(0x00000000)
	QAccessible__TitleBar             QAccessible__Role = QAccessible__Role(0x00000001)
	QAccessible__MenuBar              QAccessible__Role = QAccessible__Role(0x00000002)
	QAccessible__ScrollBar            QAccessible__Role = QAccessible__Role(0x00000003)
	QAccessible__Grip                 QAccessible__Role = QAccessible__Role(0x00000004)
	QAccessible__Sound                QAccessible__Role = QAccessible__Role(0x00000005)
	QAccessible__Cursor               QAccessible__Role = QAccessible__Role(0x00000006)
	QAccessible__Caret                QAccessible__Role = QAccessible__Role(0x00000007)
	QAccessible__AlertMessage         QAccessible__Role = QAccessible__Role(0x00000008)
	QAccessible__Window               QAccessible__Role = QAccessible__Role(0x00000009)
	QAccessible__Client               QAccessible__Role = QAccessible__Role(0x0000000A)
	QAccessible__PopupMenu            QAccessible__Role = QAccessible__Role(0x0000000B)
	QAccessible__MenuItem             QAccessible__Role = QAccessible__Role(0x0000000C)
	QAccessible__ToolTip              QAccessible__Role = QAccessible__Role(0x0000000D)
	QAccessible__Application          QAccessible__Role = QAccessible__Role(0x0000000E)
	QAccessible__Document             QAccessible__Role = QAccessible__Role(0x0000000F)
	QAccessible__Pane                 QAccessible__Role = QAccessible__Role(0x00000010)
	QAccessible__Chart                QAccessible__Role = QAccessible__Role(0x00000011)
	QAccessible__Dialog               QAccessible__Role = QAccessible__Role(0x00000012)
	QAccessible__Border               QAccessible__Role = QAccessible__Role(0x00000013)
	QAccessible__Grouping             QAccessible__Role = QAccessible__Role(0x00000014)
	QAccessible__Separator            QAccessible__Role = QAccessible__Role(0x00000015)
	QAccessible__ToolBar              QAccessible__Role = QAccessible__Role(0x00000016)
	QAccessible__StatusBar            QAccessible__Role = QAccessible__Role(0x00000017)
	QAccessible__Table                QAccessible__Role = QAccessible__Role(0x00000018)
	QAccessible__ColumnHeader         QAccessible__Role = QAccessible__Role(0x00000019)
	QAccessible__RowHeader            QAccessible__Role = QAccessible__Role(0x0000001A)
	QAccessible__Column               QAccessible__Role = QAccessible__Role(0x0000001B)
	QAccessible__Row                  QAccessible__Role = QAccessible__Role(0x0000001C)
	QAccessible__Cell                 QAccessible__Role = QAccessible__Role(0x0000001D)
	QAccessible__Link                 QAccessible__Role = QAccessible__Role(0x0000001E)
	QAccessible__HelpBalloon          QAccessible__Role = QAccessible__Role(0x0000001F)
	QAccessible__Assistant            QAccessible__Role = QAccessible__Role(0x00000020)
	QAccessible__List                 QAccessible__Role = QAccessible__Role(0x00000021)
	QAccessible__ListItem             QAccessible__Role = QAccessible__Role(0x00000022)
	QAccessible__Tree                 QAccessible__Role = QAccessible__Role(0x00000023)
	QAccessible__TreeItem             QAccessible__Role = QAccessible__Role(0x00000024)
	QAccessible__PageTab              QAccessible__Role = QAccessible__Role(0x00000025)
	QAccessible__PropertyPage         QAccessible__Role = QAccessible__Role(0x00000026)
	QAccessible__Indicator            QAccessible__Role = QAccessible__Role(0x00000027)
	QAccessible__Graphic              QAccessible__Role = QAccessible__Role(0x00000028)
	QAccessible__StaticText           QAccessible__Role = QAccessible__Role(0x00000029)
	QAccessible__EditableText         QAccessible__Role = QAccessible__Role(0x0000002A)
	QAccessible__Button               QAccessible__Role = QAccessible__Role(0x0000002B)
	QAccessible__CheckBox             QAccessible__Role = QAccessible__Role(0x0000002C)
	QAccessible__RadioButton          QAccessible__Role = QAccessible__Role(0x0000002D)
	QAccessible__ComboBox             QAccessible__Role = QAccessible__Role(0x0000002E)
	QAccessible__ProgressBar          QAccessible__Role = QAccessible__Role(0x00000030)
	QAccessible__Dial                 QAccessible__Role = QAccessible__Role(0x00000031)
	QAccessible__HotkeyField          QAccessible__Role = QAccessible__Role(0x00000032)
	QAccessible__Slider               QAccessible__Role = QAccessible__Role(0x00000033)
	QAccessible__SpinBox              QAccessible__Role = QAccessible__Role(0x00000034)
	QAccessible__Canvas               QAccessible__Role = QAccessible__Role(0x00000035)
	QAccessible__Animation            QAccessible__Role = QAccessible__Role(0x00000036)
	QAccessible__Equation             QAccessible__Role = QAccessible__Role(0x00000037)
	QAccessible__ButtonDropDown       QAccessible__Role = QAccessible__Role(0x00000038)
	QAccessible__ButtonMenu           QAccessible__Role = QAccessible__Role(0x00000039)
	QAccessible__ButtonDropGrid       QAccessible__Role = QAccessible__Role(0x0000003A)
	QAccessible__Whitespace           QAccessible__Role = QAccessible__Role(0x0000003B)
	QAccessible__PageTabList          QAccessible__Role = QAccessible__Role(0x0000003C)
	QAccessible__Clock                QAccessible__Role = QAccessible__Role(0x0000003D)
	QAccessible__Splitter             QAccessible__Role = QAccessible__Role(0x0000003E)
	QAccessible__LayeredPane          QAccessible__Role = QAccessible__Role(0x00000080)
	QAccessible__Terminal             QAccessible__Role = QAccessible__Role(0x00000081)
	QAccessible__Desktop              QAccessible__Role = QAccessible__Role(0x00000082)
	QAccessible__Paragraph            QAccessible__Role = QAccessible__Role(0x00000083)
	QAccessible__WebDocument          QAccessible__Role = QAccessible__Role(0x00000084)
	QAccessible__Section              QAccessible__Role = QAccessible__Role(0x00000085)
	QAccessible__ColorChooser         QAccessible__Role = QAccessible__Role(0x404)
	QAccessible__Footer               QAccessible__Role = QAccessible__Role(0x40E)
	QAccessible__Form                 QAccessible__Role = QAccessible__Role(0x410)
	QAccessible__Heading              QAccessible__Role = QAccessible__Role(0x414)
	QAccessible__Note                 QAccessible__Role = QAccessible__Role(0x41B)
	QAccessible__ComplementaryContent QAccessible__Role = QAccessible__Role(0x42C)
	QAccessible__UserRole             QAccessible__Role = QAccessible__Role(0x0000ffff)
)

//go:generate stringer -type=QAccessible__Text
//QAccessible::Text
type QAccessible__Text int64

const (
	QAccessible__Name             QAccessible__Text = QAccessible__Text(0)
	QAccessible__Description      QAccessible__Text = QAccessible__Text(1)
	QAccessible__Value            QAccessible__Text = QAccessible__Text(2)
	QAccessible__Help             QAccessible__Text = QAccessible__Text(3)
	QAccessible__Accelerator      QAccessible__Text = QAccessible__Text(4)
	QAccessible__DebugDescription QAccessible__Text = QAccessible__Text(5)
	QAccessible__UserText         QAccessible__Text = QAccessible__Text(0x0000ffff)
)

//go:generate stringer -type=QAccessible__TextBoundaryType
//QAccessible::TextBoundaryType
type QAccessible__TextBoundaryType int64

const (
	QAccessible__CharBoundary      QAccessible__TextBoundaryType = QAccessible__TextBoundaryType(0)
	QAccessible__WordBoundary      QAccessible__TextBoundaryType = QAccessible__TextBoundaryType(1)
	QAccessible__SentenceBoundary  QAccessible__TextBoundaryType = QAccessible__TextBoundaryType(2)
	QAccessible__ParagraphBoundary QAccessible__TextBoundaryType = QAccessible__TextBoundaryType(3)
	QAccessible__LineBoundary      QAccessible__TextBoundaryType = QAccessible__TextBoundaryType(4)
	QAccessible__NoBoundary        QAccessible__TextBoundaryType = QAccessible__TextBoundaryType(5)
)

//go:generate stringer -type=QAccessibleTableModelChangeEvent__ModelChangeType
//QAccessibleTableModelChangeEvent::ModelChangeType
type QAccessibleTableModelChangeEvent__ModelChangeType int64

const (
	QAccessibleTableModelChangeEvent__ModelReset      QAccessibleTableModelChangeEvent__ModelChangeType = QAccessibleTableModelChangeEvent__ModelChangeType(0)
	QAccessibleTableModelChangeEvent__DataChanged     QAccessibleTableModelChangeEvent__ModelChangeType = QAccessibleTableModelChangeEvent__ModelChangeType(1)
	QAccessibleTableModelChangeEvent__RowsInserted    QAccessibleTableModelChangeEvent__ModelChangeType = QAccessibleTableModelChangeEvent__ModelChangeType(2)
	QAccessibleTableModelChangeEvent__ColumnsInserted QAccessibleTableModelChangeEvent__ModelChangeType = QAccessibleTableModelChangeEvent__ModelChangeType(3)
	QAccessibleTableModelChangeEvent__RowsRemoved     QAccessibleTableModelChangeEvent__ModelChangeType = QAccessibleTableModelChangeEvent__ModelChangeType(4)
	QAccessibleTableModelChangeEvent__ColumnsRemoved  QAccessibleTableModelChangeEvent__ModelChangeType = QAccessibleTableModelChangeEvent__ModelChangeType(5)
)

type QBitmap struct {
	QPixmap
}

type QBitmap_ITF interface {
	QPixmap_ITF
	QBitmap_PTR() *QBitmap
}

func (ptr *QBitmap) QBitmap_PTR() *QBitmap {
	return ptr
}

func (ptr *QBitmap) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPixmap_PTR().Pointer()
	}
	return nil
}

func (ptr *QBitmap) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPixmap_PTR().SetPointer(p)
	}
}

func PointerFromQBitmap(ptr QBitmap_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBitmap_PTR().Pointer()
	}
	return nil
}

func NewQBitmapFromPointer(ptr unsafe.Pointer) (n *QBitmap) {
	n = new(QBitmap)
	n.SetPointer(ptr)
	return
}
func NewQBitmap() *QBitmap {
	return NewQBitmapFromPointer(C.QBitmap_NewQBitmap())
}

func NewQBitmap3(width int, height int) *QBitmap {
	return NewQBitmapFromPointer(C.QBitmap_NewQBitmap3(C.int(int32(width)), C.int(int32(height))))
}

func NewQBitmap6(other QBitmap_ITF) *QBitmap {
	return NewQBitmapFromPointer(C.QBitmap_NewQBitmap6(PointerFromQBitmap(other)))
}

func NewQBitmap2(pixmap QPixmap_ITF) *QBitmap {
	return NewQBitmapFromPointer(C.QBitmap_NewQBitmap2(PointerFromQPixmap(pixmap)))
}

func NewQBitmap4(size core.QSize_ITF) *QBitmap {
	return NewQBitmapFromPointer(C.QBitmap_NewQBitmap4(core.PointerFromQSize(size)))
}

func NewQBitmap5(fileName string, format string) *QBitmap {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	var formatC *C.char
	if format != "" {
		formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
	}
	return NewQBitmapFromPointer(C.QBitmap_NewQBitmap5(C.struct_QtGui_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, formatC))
}

func (ptr *QBitmap) DestroyQBitmap() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBitmap_DestroyQBitmap(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QClipboard__Mode
//QClipboard::Mode
type QClipboard__Mode int64

const (
	QClipboard__Clipboard  QClipboard__Mode = QClipboard__Mode(0)
	QClipboard__Selection  QClipboard__Mode = QClipboard__Mode(1)
	QClipboard__FindBuffer QClipboard__Mode = QClipboard__Mode(2)
	QClipboard__LastMode   QClipboard__Mode = QClipboard__Mode(QClipboard__FindBuffer)
)

//go:generate stringer -type=QColor__Spec
//QColor::Spec
type QColor__Spec int64

const (
	QColor__Invalid QColor__Spec = QColor__Spec(0)
	QColor__Rgb     QColor__Spec = QColor__Spec(1)
	QColor__Hsv     QColor__Spec = QColor__Spec(2)
	QColor__Cmyk    QColor__Spec = QColor__Spec(3)
	QColor__Hsl     QColor__Spec = QColor__Spec(4)
)

//go:generate stringer -type=QColor__NameFormat
//QColor::NameFormat
type QColor__NameFormat int64

const (
	QColor__HexRgb  QColor__NameFormat = QColor__NameFormat(0)
	QColor__HexArgb QColor__NameFormat = QColor__NameFormat(1)
)

//go:generate stringer -type=QContextMenuEvent__Reason
//QContextMenuEvent::Reason
type QContextMenuEvent__Reason int64

const (
	QContextMenuEvent__Mouse    QContextMenuEvent__Reason = QContextMenuEvent__Reason(0)
	QContextMenuEvent__Keyboard QContextMenuEvent__Reason = QContextMenuEvent__Reason(1)
	QContextMenuEvent__Other    QContextMenuEvent__Reason = QContextMenuEvent__Reason(2)
)

//go:generate stringer -type=QDoubleValidator__Notation
//QDoubleValidator::Notation
type QDoubleValidator__Notation int64

const (
	QDoubleValidator__StandardNotation   QDoubleValidator__Notation = QDoubleValidator__Notation(0)
	QDoubleValidator__ScientificNotation QDoubleValidator__Notation = QDoubleValidator__Notation(1)
)

//go:generate stringer -type=QFont__Capitalization
//QFont::Capitalization
type QFont__Capitalization int64

const (
	QFont__MixedCase    QFont__Capitalization = QFont__Capitalization(0)
	QFont__AllUppercase QFont__Capitalization = QFont__Capitalization(1)
	QFont__AllLowercase QFont__Capitalization = QFont__Capitalization(2)
	QFont__SmallCaps    QFont__Capitalization = QFont__Capitalization(3)
	QFont__Capitalize   QFont__Capitalization = QFont__Capitalization(4)
)

//go:generate stringer -type=QFont__HintingPreference
//QFont::HintingPreference
type QFont__HintingPreference int64

const (
	QFont__PreferDefaultHinting  QFont__HintingPreference = QFont__HintingPreference(0)
	QFont__PreferNoHinting       QFont__HintingPreference = QFont__HintingPreference(1)
	QFont__PreferVerticalHinting QFont__HintingPreference = QFont__HintingPreference(2)
	QFont__PreferFullHinting     QFont__HintingPreference = QFont__HintingPreference(3)
)

//go:generate stringer -type=QFont__SpacingType
//QFont::SpacingType
type QFont__SpacingType int64

const (
	QFont__PercentageSpacing QFont__SpacingType = QFont__SpacingType(0)
	QFont__AbsoluteSpacing   QFont__SpacingType = QFont__SpacingType(1)
)

//go:generate stringer -type=QFont__Stretch
//QFont::Stretch
type QFont__Stretch int64

const (
	QFont__AnyStretch     QFont__Stretch = QFont__Stretch(0)
	QFont__UltraCondensed QFont__Stretch = QFont__Stretch(50)
	QFont__ExtraCondensed QFont__Stretch = QFont__Stretch(62)
	QFont__Condensed      QFont__Stretch = QFont__Stretch(75)
	QFont__SemiCondensed  QFont__Stretch = QFont__Stretch(87)
	QFont__Unstretched    QFont__Stretch = QFont__Stretch(100)
	QFont__SemiExpanded   QFont__Stretch = QFont__Stretch(112)
	QFont__Expanded       QFont__Stretch = QFont__Stretch(125)
	QFont__ExtraExpanded  QFont__Stretch = QFont__Stretch(150)
	QFont__UltraExpanded  QFont__Stretch = QFont__Stretch(200)
)

//go:generate stringer -type=QFont__Style
//QFont::Style
type QFont__Style int64

var (
	QFont__StyleNormal  QFont__Style = QFont__Style(0)
	QFont__StyleItalic  QFont__Style = QFont__Style(1)
	QFont__StyleOblique QFont__Style = QFont__Style(2)
)

//go:generate stringer -type=QFont__StyleHint
//QFont::StyleHint
type QFont__StyleHint int64

var (
	QFont__Helvetica  QFont__StyleHint = QFont__StyleHint(0)
	QFont__SansSerif  QFont__StyleHint = QFont__StyleHint(QFont__Helvetica)
	QFont__Times      QFont__StyleHint = QFont__StyleHint(C.QFont_Times_Type())
	QFont__Serif      QFont__StyleHint = QFont__StyleHint(QFont__Times)
	QFont__Courier    QFont__StyleHint = QFont__StyleHint(C.QFont_Courier_Type())
	QFont__TypeWriter QFont__StyleHint = QFont__StyleHint(QFont__Courier)
	QFont__OldEnglish QFont__StyleHint = QFont__StyleHint(C.QFont_OldEnglish_Type())
	QFont__Decorative QFont__StyleHint = QFont__StyleHint(QFont__OldEnglish)
	QFont__System     QFont__StyleHint = QFont__StyleHint(C.QFont_System_Type())
	QFont__AnyStyle   QFont__StyleHint = QFont__StyleHint(C.QFont_AnyStyle_Type())
	QFont__Cursive    QFont__StyleHint = QFont__StyleHint(C.QFont_Cursive_Type())
	QFont__Monospace  QFont__StyleHint = QFont__StyleHint(C.QFont_Monospace_Type())
	QFont__Fantasy    QFont__StyleHint = QFont__StyleHint(C.QFont_Fantasy_Type())
)

//go:generate stringer -type=QFont__StyleStrategy
//QFont::StyleStrategy
type QFont__StyleStrategy int64

var (
	QFont__PreferDefault       QFont__StyleStrategy = QFont__StyleStrategy(0x0001)
	QFont__PreferBitmap        QFont__StyleStrategy = QFont__StyleStrategy(0x0002)
	QFont__PreferDevice        QFont__StyleStrategy = QFont__StyleStrategy(0x0004)
	QFont__PreferOutline       QFont__StyleStrategy = QFont__StyleStrategy(0x0008)
	QFont__ForceOutline        QFont__StyleStrategy = QFont__StyleStrategy(0x0010)
	QFont__PreferMatch         QFont__StyleStrategy = QFont__StyleStrategy(0x0020)
	QFont__PreferQuality       QFont__StyleStrategy = QFont__StyleStrategy(0x0040)
	QFont__PreferAntialias     QFont__StyleStrategy = QFont__StyleStrategy(0x0080)
	QFont__NoAntialias         QFont__StyleStrategy = QFont__StyleStrategy(0x0100)
	QFont__OpenGLCompatible    QFont__StyleStrategy = QFont__StyleStrategy(0x0200)
	QFont__ForceIntegerMetrics QFont__StyleStrategy = QFont__StyleStrategy(0x0400)
	QFont__NoSubpixelAntialias QFont__StyleStrategy = QFont__StyleStrategy(0x0800)
	QFont__NoFontMerging       QFont__StyleStrategy = QFont__StyleStrategy(0x8000)
)

//go:generate stringer -type=QFont__Weight
//QFont::Weight
type QFont__Weight int64

const (
	QFont__Thin       QFont__Weight = QFont__Weight(0)
	QFont__ExtraLight QFont__Weight = QFont__Weight(12)
	QFont__Light      QFont__Weight = QFont__Weight(25)
	QFont__Normal     QFont__Weight = QFont__Weight(50)
	QFont__Medium     QFont__Weight = QFont__Weight(57)
	QFont__DemiBold   QFont__Weight = QFont__Weight(63)
	QFont__Bold       QFont__Weight = QFont__Weight(75)
	QFont__ExtraBold  QFont__Weight = QFont__Weight(81)
	QFont__Black      QFont__Weight = QFont__Weight(87)
)

//go:generate stringer -type=QFontDatabase__WritingSystem
//QFontDatabase::WritingSystem
type QFontDatabase__WritingSystem int64

var (
	QFontDatabase__Any                 QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(0)
	QFontDatabase__Latin               QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(1)
	QFontDatabase__Greek               QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(2)
	QFontDatabase__Cyrillic            QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(3)
	QFontDatabase__Armenian            QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(4)
	QFontDatabase__Hebrew              QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(5)
	QFontDatabase__Arabic              QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(6)
	QFontDatabase__Syriac              QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(7)
	QFontDatabase__Thaana              QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(8)
	QFontDatabase__Devanagari          QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(9)
	QFontDatabase__Bengali             QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(10)
	QFontDatabase__Gurmukhi            QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(11)
	QFontDatabase__Gujarati            QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(12)
	QFontDatabase__Oriya               QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(13)
	QFontDatabase__Tamil               QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(14)
	QFontDatabase__Telugu              QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(15)
	QFontDatabase__Kannada             QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(16)
	QFontDatabase__Malayalam           QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(17)
	QFontDatabase__Sinhala             QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(18)
	QFontDatabase__Thai                QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(19)
	QFontDatabase__Lao                 QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(20)
	QFontDatabase__Tibetan             QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(21)
	QFontDatabase__Myanmar             QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(22)
	QFontDatabase__Georgian            QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(23)
	QFontDatabase__Khmer               QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(24)
	QFontDatabase__SimplifiedChinese   QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(25)
	QFontDatabase__TraditionalChinese  QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(26)
	QFontDatabase__Japanese            QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(27)
	QFontDatabase__Korean              QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(28)
	QFontDatabase__Vietnamese          QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(29)
	QFontDatabase__Symbol              QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(30)
	QFontDatabase__Other               QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(QFontDatabase__Symbol)
	QFontDatabase__Ogham               QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(C.QFontDatabase_Ogham_Type())
	QFontDatabase__Runic               QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(C.QFontDatabase_Runic_Type())
	QFontDatabase__Nko                 QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(C.QFontDatabase_Nko_Type())
	QFontDatabase__WritingSystemsCount QFontDatabase__WritingSystem = QFontDatabase__WritingSystem(C.QFontDatabase_WritingSystemsCount_Type())
)

//go:generate stringer -type=QFontDatabase__SystemFont
//QFontDatabase::SystemFont
type QFontDatabase__SystemFont int64

const (
	QFontDatabase__GeneralFont          QFontDatabase__SystemFont = QFontDatabase__SystemFont(0)
	QFontDatabase__FixedFont            QFontDatabase__SystemFont = QFontDatabase__SystemFont(1)
	QFontDatabase__TitleFont            QFontDatabase__SystemFont = QFontDatabase__SystemFont(2)
	QFontDatabase__SmallestReadableFont QFontDatabase__SystemFont = QFontDatabase__SystemFont(3)
)

//go:generate stringer -type=QGlyphRun__GlyphRunFlag
//QGlyphRun::GlyphRunFlag
type QGlyphRun__GlyphRunFlag int64

const (
	QGlyphRun__Overline      QGlyphRun__GlyphRunFlag = QGlyphRun__GlyphRunFlag(0x01)
	QGlyphRun__Underline     QGlyphRun__GlyphRunFlag = QGlyphRun__GlyphRunFlag(0x02)
	QGlyphRun__StrikeOut     QGlyphRun__GlyphRunFlag = QGlyphRun__GlyphRunFlag(0x04)
	QGlyphRun__RightToLeft   QGlyphRun__GlyphRunFlag = QGlyphRun__GlyphRunFlag(0x08)
	QGlyphRun__SplitLigature QGlyphRun__GlyphRunFlag = QGlyphRun__GlyphRunFlag(0x10)
)

//go:generate stringer -type=QGradient__Type
//QGradient::Type
type QGradient__Type int64

const (
	QGradient__LinearGradient  QGradient__Type = QGradient__Type(0)
	QGradient__RadialGradient  QGradient__Type = QGradient__Type(1)
	QGradient__ConicalGradient QGradient__Type = QGradient__Type(2)
	QGradient__NoGradient      QGradient__Type = QGradient__Type(3)
)

//go:generate stringer -type=QGradient__CoordinateMode
//QGradient::CoordinateMode
type QGradient__CoordinateMode int64

const (
	QGradient__LogicalMode         QGradient__CoordinateMode = QGradient__CoordinateMode(0)
	QGradient__StretchToDeviceMode QGradient__CoordinateMode = QGradient__CoordinateMode(1)
	QGradient__ObjectBoundingMode  QGradient__CoordinateMode = QGradient__CoordinateMode(2)
)

//go:generate stringer -type=QGradient__Spread
//QGradient::Spread
type QGradient__Spread int64

const (
	QGradient__PadSpread     QGradient__Spread = QGradient__Spread(0)
	QGradient__ReflectSpread QGradient__Spread = QGradient__Spread(1)
	QGradient__RepeatSpread  QGradient__Spread = QGradient__Spread(2)
)

type QGuiApplication struct {
	core.QCoreApplication
}

type QGuiApplication_ITF interface {
	core.QCoreApplication_ITF
	QGuiApplication_PTR() *QGuiApplication
}

func (ptr *QGuiApplication) QGuiApplication_PTR() *QGuiApplication {
	return ptr
}

func (ptr *QGuiApplication) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QCoreApplication_PTR().Pointer()
	}
	return nil
}

func (ptr *QGuiApplication) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QCoreApplication_PTR().SetPointer(p)
	}
}

func PointerFromQGuiApplication(ptr QGuiApplication_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGuiApplication_PTR().Pointer()
	}
	return nil
}

func NewQGuiApplicationFromPointer(ptr unsafe.Pointer) (n *QGuiApplication) {
	n = new(QGuiApplication)
	n.SetPointer(ptr)
	return
}
func NewQGuiApplication(argc int, argv []string) *QGuiApplication {
	argvC := C.CString(strings.Join(argv, "|"))
	defer C.free(unsafe.Pointer(argvC))
	tmpValue := NewQGuiApplicationFromPointer(C.QGuiApplication_NewQGuiApplication(C.int(int32(argc)), argvC))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QGuiApplication_Exec() int {
	return int(int32(C.QGuiApplication_QGuiApplication_Exec()))
}

func (ptr *QGuiApplication) Exec() int {
	return int(int32(C.QGuiApplication_QGuiApplication_Exec()))
}

//export callbackQGuiApplication_DestroyQGuiApplication
func callbackQGuiApplication_DestroyQGuiApplication(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QGuiApplication"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQGuiApplicationFromPointer(ptr).DestroyQGuiApplicationDefault()
	}
}

func (ptr *QGuiApplication) ConnectDestroyQGuiApplication(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QGuiApplication"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QGuiApplication", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QGuiApplication", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QGuiApplication) DisconnectDestroyQGuiApplication() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QGuiApplication")
	}
}

func (ptr *QGuiApplication) DestroyQGuiApplication() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGuiApplication_DestroyQGuiApplication(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGuiApplication) DestroyQGuiApplicationDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QGuiApplication_DestroyQGuiApplicationDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGuiApplication) __screens_atList(i int) *QScreen {
	if ptr.Pointer() != nil {
		tmpValue := NewQScreenFromPointer(C.QGuiApplication___screens_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGuiApplication) __screens_setList(i QScreen_ITF) {
	if ptr.Pointer() != nil {
		C.QGuiApplication___screens_setList(ptr.Pointer(), PointerFromQScreen(i))
	}
}

func (ptr *QGuiApplication) __screens_newList() unsafe.Pointer {
	return C.QGuiApplication___screens_newList(ptr.Pointer())
}

func (ptr *QGuiApplication) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QGuiApplication___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QGuiApplication) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QGuiApplication___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QGuiApplication) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QGuiApplication___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QGuiApplication) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGuiApplication___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGuiApplication) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGuiApplication___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGuiApplication) __findChildren_newList2() unsafe.Pointer {
	return C.QGuiApplication___findChildren_newList2(ptr.Pointer())
}

func (ptr *QGuiApplication) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGuiApplication___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGuiApplication) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGuiApplication___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGuiApplication) __findChildren_newList3() unsafe.Pointer {
	return C.QGuiApplication___findChildren_newList3(ptr.Pointer())
}

func (ptr *QGuiApplication) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGuiApplication___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGuiApplication) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGuiApplication___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGuiApplication) __findChildren_newList() unsafe.Pointer {
	return C.QGuiApplication___findChildren_newList(ptr.Pointer())
}

func (ptr *QGuiApplication) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QGuiApplication___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGuiApplication) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGuiApplication___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QGuiApplication) __children_newList() unsafe.Pointer {
	return C.QGuiApplication___children_newList(ptr.Pointer())
}

//export callbackQGuiApplication_ObjectNameChanged
func callbackQGuiApplication_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtGui_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

type QIcon struct {
	ptr unsafe.Pointer
}

type QIcon_ITF interface {
	QIcon_PTR() *QIcon
}

func (ptr *QIcon) QIcon_PTR() *QIcon {
	return ptr
}

func (ptr *QIcon) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QIcon) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQIcon(ptr QIcon_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIcon_PTR().Pointer()
	}
	return nil
}

func NewQIconFromPointer(ptr unsafe.Pointer) (n *QIcon) {
	n = new(QIcon)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QIcon__Mode
//QIcon::Mode
type QIcon__Mode int64

const (
	QIcon__Normal   QIcon__Mode = QIcon__Mode(0)
	QIcon__Disabled QIcon__Mode = QIcon__Mode(1)
	QIcon__Active   QIcon__Mode = QIcon__Mode(2)
	QIcon__Selected QIcon__Mode = QIcon__Mode(3)
)

//go:generate stringer -type=QIcon__State
//QIcon::State
type QIcon__State int64

const (
	QIcon__On  QIcon__State = QIcon__State(0)
	QIcon__Off QIcon__State = QIcon__State(1)
)

func NewQIcon() *QIcon {
	tmpValue := NewQIconFromPointer(C.QIcon_NewQIcon())
	qt.SetFinalizer(tmpValue, (*QIcon).DestroyQIcon)
	return tmpValue
}

func NewQIcon4(other QIcon_ITF) *QIcon {
	tmpValue := NewQIconFromPointer(C.QIcon_NewQIcon4(PointerFromQIcon(other)))
	qt.SetFinalizer(tmpValue, (*QIcon).DestroyQIcon)
	return tmpValue
}

func NewQIcon6(engine QIconEngine_ITF) *QIcon {
	tmpValue := NewQIconFromPointer(C.QIcon_NewQIcon6(PointerFromQIconEngine(engine)))
	qt.SetFinalizer(tmpValue, (*QIcon).DestroyQIcon)
	return tmpValue
}

func NewQIcon3(other QIcon_ITF) *QIcon {
	tmpValue := NewQIconFromPointer(C.QIcon_NewQIcon3(PointerFromQIcon(other)))
	qt.SetFinalizer(tmpValue, (*QIcon).DestroyQIcon)
	return tmpValue
}

func NewQIcon2(pixmap QPixmap_ITF) *QIcon {
	tmpValue := NewQIconFromPointer(C.QIcon_NewQIcon2(PointerFromQPixmap(pixmap)))
	qt.SetFinalizer(tmpValue, (*QIcon).DestroyQIcon)
	return tmpValue
}

func NewQIcon5(fileName string) *QIcon {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	tmpValue := NewQIconFromPointer(C.QIcon_NewQIcon5(C.struct_QtGui_PackedString{data: fileNameC, len: C.longlong(len(fileName))}))
	qt.SetFinalizer(tmpValue, (*QIcon).DestroyQIcon)
	return tmpValue
}

func (ptr *QIcon) DestroyQIcon() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QIcon_DestroyQIcon(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QIcon) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QIcon_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QIcon) ToVariant() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QIcon_ToVariant(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QIcon) __availableSizes_atList(i int) *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QIcon___availableSizes_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QIcon) __availableSizes_setList(i core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QIcon___availableSizes_setList(ptr.Pointer(), core.PointerFromQSize(i))
	}
}

func (ptr *QIcon) __availableSizes_newList() unsafe.Pointer {
	return C.QIcon___availableSizes_newList(ptr.Pointer())
}

type QIconEngine struct {
	ptr unsafe.Pointer
}

type QIconEngine_ITF interface {
	QIconEngine_PTR() *QIconEngine
}

func (ptr *QIconEngine) QIconEngine_PTR() *QIconEngine {
	return ptr
}

func (ptr *QIconEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QIconEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQIconEngine(ptr QIconEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIconEngine_PTR().Pointer()
	}
	return nil
}

func NewQIconEngineFromPointer(ptr unsafe.Pointer) (n *QIconEngine) {
	n = new(QIconEngine)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QIconEngine__IconEngineHook
//QIconEngine::IconEngineHook
type QIconEngine__IconEngineHook int64

const (
	QIconEngine__AvailableSizesHook QIconEngine__IconEngineHook = QIconEngine__IconEngineHook(1)
	QIconEngine__IconNameHook       QIconEngine__IconEngineHook = QIconEngine__IconEngineHook(2)
	QIconEngine__IsNullHook         QIconEngine__IconEngineHook = QIconEngine__IconEngineHook(3)
	QIconEngine__ScaledPixmapHook   QIconEngine__IconEngineHook = QIconEngine__IconEngineHook(4)
)

func NewQIconEngine() *QIconEngine {
	return NewQIconEngineFromPointer(C.QIconEngine_NewQIconEngine())
}

//export callbackQIconEngine_Read
func callbackQIconEngine_Read(ptr unsafe.Pointer, in unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "read"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QDataStream) bool)(signal))(core.NewQDataStreamFromPointer(in)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQIconEngineFromPointer(ptr).ReadDefault(core.NewQDataStreamFromPointer(in)))))
}

func (ptr *QIconEngine) ConnectRead(f func(in *core.QDataStream) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "read"); signal != nil {
			f := func(in *core.QDataStream) bool {
				(*(*func(*core.QDataStream) bool)(signal))(in)
				return f(in)
			}
			qt.ConnectSignal(ptr.Pointer(), "read", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "read", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIconEngine) DisconnectRead() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "read")
	}
}

func (ptr *QIconEngine) Read(in core.QDataStream_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QIconEngine_Read(ptr.Pointer(), core.PointerFromQDataStream(in))) != 0
	}
	return false
}

func (ptr *QIconEngine) ReadDefault(in core.QDataStream_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QIconEngine_ReadDefault(ptr.Pointer(), core.PointerFromQDataStream(in))) != 0
	}
	return false
}

//export callbackQIconEngine_Paint
func callbackQIconEngine_Paint(ptr unsafe.Pointer, painter unsafe.Pointer, rect unsafe.Pointer, mode C.longlong, state C.longlong) {
	if signal := qt.GetSignal(ptr, "paint"); signal != nil {
		(*(*func(*QPainter, *core.QRect, QIcon__Mode, QIcon__State))(signal))(NewQPainterFromPointer(painter), core.NewQRectFromPointer(rect), QIcon__Mode(mode), QIcon__State(state))
	}

}

func (ptr *QIconEngine) ConnectPaint(f func(painter *QPainter, rect *core.QRect, mode QIcon__Mode, state QIcon__State)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "paint"); signal != nil {
			f := func(painter *QPainter, rect *core.QRect, mode QIcon__Mode, state QIcon__State) {
				(*(*func(*QPainter, *core.QRect, QIcon__Mode, QIcon__State))(signal))(painter, rect, mode, state)
				f(painter, rect, mode, state)
			}
			qt.ConnectSignal(ptr.Pointer(), "paint", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "paint", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIconEngine) DisconnectPaint() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "paint")
	}
}

func (ptr *QIconEngine) Paint(painter QPainter_ITF, rect core.QRect_ITF, mode QIcon__Mode, state QIcon__State) {
	if ptr.Pointer() != nil {
		C.QIconEngine_Paint(ptr.Pointer(), PointerFromQPainter(painter), core.PointerFromQRect(rect), C.longlong(mode), C.longlong(state))
	}
}

//export callbackQIconEngine_DestroyQIconEngine
func callbackQIconEngine_DestroyQIconEngine(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QIconEngine"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQIconEngineFromPointer(ptr).DestroyQIconEngineDefault()
	}
}

func (ptr *QIconEngine) ConnectDestroyQIconEngine(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QIconEngine"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QIconEngine", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QIconEngine", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIconEngine) DisconnectDestroyQIconEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QIconEngine")
	}
}

func (ptr *QIconEngine) DestroyQIconEngine() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QIconEngine_DestroyQIconEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QIconEngine) DestroyQIconEngineDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QIconEngine_DestroyQIconEngineDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQIconEngine_Clone
func callbackQIconEngine_Clone(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "clone"); signal != nil {
		return PointerFromQIconEngine((*(*func() *QIconEngine)(signal))())
	}

	return PointerFromQIconEngine(NewQIconEngine())
}

func (ptr *QIconEngine) ConnectClone(f func() *QIconEngine) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clone"); signal != nil {
			f := func() *QIconEngine {
				(*(*func() *QIconEngine)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clone", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clone", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIconEngine) DisconnectClone() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clone")
	}
}

func (ptr *QIconEngine) Clone() *QIconEngine {
	if ptr.Pointer() != nil {
		return NewQIconEngineFromPointer(C.QIconEngine_Clone(ptr.Pointer()))
	}
	return nil
}

//export callbackQIconEngine_Key
func callbackQIconEngine_Key(ptr unsafe.Pointer) C.struct_QtGui_PackedString {
	if signal := qt.GetSignal(ptr, "key"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_QtGui_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewQIconEngineFromPointer(ptr).KeyDefault()
	return C.struct_QtGui_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QIconEngine) ConnectKey(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "key"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "key", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "key", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIconEngine) DisconnectKey() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "key")
	}
}

func (ptr *QIconEngine) Key() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QIconEngine_Key(ptr.Pointer()))
	}
	return ""
}

func (ptr *QIconEngine) KeyDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QIconEngine_KeyDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQIconEngine_Write
func callbackQIconEngine_Write(ptr unsafe.Pointer, out unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "write"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QDataStream) bool)(signal))(core.NewQDataStreamFromPointer(out)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQIconEngineFromPointer(ptr).WriteDefault(core.NewQDataStreamFromPointer(out)))))
}

func (ptr *QIconEngine) ConnectWrite(f func(out *core.QDataStream) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "write"); signal != nil {
			f := func(out *core.QDataStream) bool {
				(*(*func(*core.QDataStream) bool)(signal))(out)
				return f(out)
			}
			qt.ConnectSignal(ptr.Pointer(), "write", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "write", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QIconEngine) DisconnectWrite() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "write")
	}
}

func (ptr *QIconEngine) Write(out core.QDataStream_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QIconEngine_Write(ptr.Pointer(), core.PointerFromQDataStream(out))) != 0
	}
	return false
}

func (ptr *QIconEngine) WriteDefault(out core.QDataStream_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QIconEngine_WriteDefault(ptr.Pointer(), core.PointerFromQDataStream(out))) != 0
	}
	return false
}

func (ptr *QIconEngine) __availableSizes_atList(i int) *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QIconEngine___availableSizes_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QIconEngine) __availableSizes_setList(i core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QIconEngine___availableSizes_setList(ptr.Pointer(), core.PointerFromQSize(i))
	}
}

func (ptr *QIconEngine) __availableSizes_newList() unsafe.Pointer {
	return C.QIconEngine___availableSizes_newList(ptr.Pointer())
}

//go:generate stringer -type=QImage__Format
//QImage::Format
type QImage__Format int64

const (
	QImage__Format_Invalid                QImage__Format = QImage__Format(0)
	QImage__Format_Mono                   QImage__Format = QImage__Format(1)
	QImage__Format_MonoLSB                QImage__Format = QImage__Format(2)
	QImage__Format_Indexed8               QImage__Format = QImage__Format(3)
	QImage__Format_RGB32                  QImage__Format = QImage__Format(4)
	QImage__Format_ARGB32                 QImage__Format = QImage__Format(5)
	QImage__Format_ARGB32_Premultiplied   QImage__Format = QImage__Format(6)
	QImage__Format_RGB16                  QImage__Format = QImage__Format(7)
	QImage__Format_ARGB8565_Premultiplied QImage__Format = QImage__Format(8)
	QImage__Format_RGB666                 QImage__Format = QImage__Format(9)
	QImage__Format_ARGB6666_Premultiplied QImage__Format = QImage__Format(10)
	QImage__Format_RGB555                 QImage__Format = QImage__Format(11)
	QImage__Format_ARGB8555_Premultiplied QImage__Format = QImage__Format(12)
	QImage__Format_RGB888                 QImage__Format = QImage__Format(13)
	QImage__Format_RGB444                 QImage__Format = QImage__Format(14)
	QImage__Format_ARGB4444_Premultiplied QImage__Format = QImage__Format(15)
	QImage__Format_RGBX8888               QImage__Format = QImage__Format(16)
	QImage__Format_RGBA8888               QImage__Format = QImage__Format(17)
	QImage__Format_RGBA8888_Premultiplied QImage__Format = QImage__Format(18)
	QImage__Format_BGR30                  QImage__Format = QImage__Format(19)
	QImage__Format_A2BGR30_Premultiplied  QImage__Format = QImage__Format(20)
	QImage__Format_RGB30                  QImage__Format = QImage__Format(21)
	QImage__Format_A2RGB30_Premultiplied  QImage__Format = QImage__Format(22)
	QImage__Format_Alpha8                 QImage__Format = QImage__Format(23)
	QImage__Format_Grayscale8             QImage__Format = QImage__Format(24)
)

//go:generate stringer -type=QImage__InvertMode
//QImage::InvertMode
type QImage__InvertMode int64

const (
	QImage__InvertRgb  QImage__InvertMode = QImage__InvertMode(0)
	QImage__InvertRgba QImage__InvertMode = QImage__InvertMode(1)
)

//go:generate stringer -type=QImageIOHandler__ImageOption
//QImageIOHandler::ImageOption
type QImageIOHandler__ImageOption int64

const (
	QImageIOHandler__Size                 QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(0)
	QImageIOHandler__ClipRect             QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(1)
	QImageIOHandler__Description          QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(2)
	QImageIOHandler__ScaledClipRect       QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(3)
	QImageIOHandler__ScaledSize           QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(4)
	QImageIOHandler__CompressionRatio     QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(5)
	QImageIOHandler__Gamma                QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(6)
	QImageIOHandler__Quality              QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(7)
	QImageIOHandler__Name                 QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(8)
	QImageIOHandler__SubType              QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(9)
	QImageIOHandler__IncrementalReading   QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(10)
	QImageIOHandler__Endianness           QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(11)
	QImageIOHandler__Animation            QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(12)
	QImageIOHandler__BackgroundColor      QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(13)
	QImageIOHandler__ImageFormat          QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(14)
	QImageIOHandler__SupportedSubTypes    QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(15)
	QImageIOHandler__OptimizedWrite       QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(16)
	QImageIOHandler__ProgressiveScanWrite QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(17)
	QImageIOHandler__ImageTransformation  QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(18)
	QImageIOHandler__TransformedByDefault QImageIOHandler__ImageOption = QImageIOHandler__ImageOption(19)
)

//go:generate stringer -type=QImageIOHandler__Transformation
//QImageIOHandler::Transformation
type QImageIOHandler__Transformation int64

const (
	QImageIOHandler__TransformationNone              QImageIOHandler__Transformation = QImageIOHandler__Transformation(0)
	QImageIOHandler__TransformationMirror            QImageIOHandler__Transformation = QImageIOHandler__Transformation(1)
	QImageIOHandler__TransformationFlip              QImageIOHandler__Transformation = QImageIOHandler__Transformation(2)
	QImageIOHandler__TransformationRotate180         QImageIOHandler__Transformation = QImageIOHandler__Transformation(QImageIOHandler__TransformationMirror | QImageIOHandler__TransformationFlip)
	QImageIOHandler__TransformationRotate90          QImageIOHandler__Transformation = QImageIOHandler__Transformation(4)
	QImageIOHandler__TransformationMirrorAndRotate90 QImageIOHandler__Transformation = QImageIOHandler__Transformation(QImageIOHandler__TransformationMirror | QImageIOHandler__TransformationRotate90)
	QImageIOHandler__TransformationFlipAndRotate90   QImageIOHandler__Transformation = QImageIOHandler__Transformation(QImageIOHandler__TransformationFlip | QImageIOHandler__TransformationRotate90)
	QImageIOHandler__TransformationRotate270         QImageIOHandler__Transformation = QImageIOHandler__Transformation(QImageIOHandler__TransformationRotate180 | QImageIOHandler__TransformationRotate90)
)

//go:generate stringer -type=QImageIOPlugin__Capability
//QImageIOPlugin::Capability
type QImageIOPlugin__Capability int64

const (
	QImageIOPlugin__CanRead            QImageIOPlugin__Capability = QImageIOPlugin__Capability(0x1)
	QImageIOPlugin__CanWrite           QImageIOPlugin__Capability = QImageIOPlugin__Capability(0x2)
	QImageIOPlugin__CanReadIncremental QImageIOPlugin__Capability = QImageIOPlugin__Capability(0x4)
)

//go:generate stringer -type=QImageReader__ImageReaderError
//QImageReader::ImageReaderError
type QImageReader__ImageReaderError int64

const (
	QImageReader__UnknownError           QImageReader__ImageReaderError = QImageReader__ImageReaderError(0)
	QImageReader__FileNotFoundError      QImageReader__ImageReaderError = QImageReader__ImageReaderError(1)
	QImageReader__DeviceError            QImageReader__ImageReaderError = QImageReader__ImageReaderError(2)
	QImageReader__UnsupportedFormatError QImageReader__ImageReaderError = QImageReader__ImageReaderError(3)
	QImageReader__InvalidDataError       QImageReader__ImageReaderError = QImageReader__ImageReaderError(4)
)

//go:generate stringer -type=QImageWriter__ImageWriterError
//QImageWriter::ImageWriterError
type QImageWriter__ImageWriterError int64

const (
	QImageWriter__UnknownError           QImageWriter__ImageWriterError = QImageWriter__ImageWriterError(0)
	QImageWriter__DeviceError            QImageWriter__ImageWriterError = QImageWriter__ImageWriterError(1)
	QImageWriter__UnsupportedFormatError QImageWriter__ImageWriterError = QImageWriter__ImageWriterError(2)
)

//go:generate stringer -type=QInputMethod__Action
//QInputMethod::Action
type QInputMethod__Action int64

const (
	QInputMethod__Click       QInputMethod__Action = QInputMethod__Action(0)
	QInputMethod__ContextMenu QInputMethod__Action = QInputMethod__Action(1)
)

//go:generate stringer -type=QInputMethodEvent__AttributeType
//QInputMethodEvent::AttributeType
type QInputMethodEvent__AttributeType int64

const (
	QInputMethodEvent__TextFormat QInputMethodEvent__AttributeType = QInputMethodEvent__AttributeType(0)
	QInputMethodEvent__Cursor     QInputMethodEvent__AttributeType = QInputMethodEvent__AttributeType(1)
	QInputMethodEvent__Language   QInputMethodEvent__AttributeType = QInputMethodEvent__AttributeType(2)
	QInputMethodEvent__Ruby       QInputMethodEvent__AttributeType = QInputMethodEvent__AttributeType(3)
	QInputMethodEvent__Selection  QInputMethodEvent__AttributeType = QInputMethodEvent__AttributeType(4)
)

type QKeySequence struct {
	ptr unsafe.Pointer
}

type QKeySequence_ITF interface {
	QKeySequence_PTR() *QKeySequence
}

func (ptr *QKeySequence) QKeySequence_PTR() *QKeySequence {
	return ptr
}

func (ptr *QKeySequence) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QKeySequence) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQKeySequence(ptr QKeySequence_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QKeySequence_PTR().Pointer()
	}
	return nil
}

func NewQKeySequenceFromPointer(ptr unsafe.Pointer) (n *QKeySequence) {
	n = new(QKeySequence)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QKeySequence__SequenceFormat
//QKeySequence::SequenceFormat
type QKeySequence__SequenceFormat int64

const (
	QKeySequence__NativeText   QKeySequence__SequenceFormat = QKeySequence__SequenceFormat(0)
	QKeySequence__PortableText QKeySequence__SequenceFormat = QKeySequence__SequenceFormat(1)
)

//go:generate stringer -type=QKeySequence__SequenceMatch
//QKeySequence::SequenceMatch
type QKeySequence__SequenceMatch int64

const (
	QKeySequence__NoMatch      QKeySequence__SequenceMatch = QKeySequence__SequenceMatch(0)
	QKeySequence__PartialMatch QKeySequence__SequenceMatch = QKeySequence__SequenceMatch(1)
	QKeySequence__ExactMatch   QKeySequence__SequenceMatch = QKeySequence__SequenceMatch(2)
)

//go:generate stringer -type=QKeySequence__StandardKey
//QKeySequence::StandardKey
type QKeySequence__StandardKey int64

const (
	QKeySequence__UnknownKey               QKeySequence__StandardKey = QKeySequence__StandardKey(0)
	QKeySequence__HelpContents             QKeySequence__StandardKey = QKeySequence__StandardKey(1)
	QKeySequence__WhatsThis                QKeySequence__StandardKey = QKeySequence__StandardKey(2)
	QKeySequence__Open                     QKeySequence__StandardKey = QKeySequence__StandardKey(3)
	QKeySequence__Close                    QKeySequence__StandardKey = QKeySequence__StandardKey(4)
	QKeySequence__Save                     QKeySequence__StandardKey = QKeySequence__StandardKey(5)
	QKeySequence__New                      QKeySequence__StandardKey = QKeySequence__StandardKey(6)
	QKeySequence__Delete                   QKeySequence__StandardKey = QKeySequence__StandardKey(7)
	QKeySequence__Cut                      QKeySequence__StandardKey = QKeySequence__StandardKey(8)
	QKeySequence__Copy                     QKeySequence__StandardKey = QKeySequence__StandardKey(9)
	QKeySequence__Paste                    QKeySequence__StandardKey = QKeySequence__StandardKey(10)
	QKeySequence__Undo                     QKeySequence__StandardKey = QKeySequence__StandardKey(11)
	QKeySequence__Redo                     QKeySequence__StandardKey = QKeySequence__StandardKey(12)
	QKeySequence__Back                     QKeySequence__StandardKey = QKeySequence__StandardKey(13)
	QKeySequence__Forward                  QKeySequence__StandardKey = QKeySequence__StandardKey(14)
	QKeySequence__Refresh                  QKeySequence__StandardKey = QKeySequence__StandardKey(15)
	QKeySequence__ZoomIn                   QKeySequence__StandardKey = QKeySequence__StandardKey(16)
	QKeySequence__ZoomOut                  QKeySequence__StandardKey = QKeySequence__StandardKey(17)
	QKeySequence__Print                    QKeySequence__StandardKey = QKeySequence__StandardKey(18)
	QKeySequence__AddTab                   QKeySequence__StandardKey = QKeySequence__StandardKey(19)
	QKeySequence__NextChild                QKeySequence__StandardKey = QKeySequence__StandardKey(20)
	QKeySequence__PreviousChild            QKeySequence__StandardKey = QKeySequence__StandardKey(21)
	QKeySequence__Find                     QKeySequence__StandardKey = QKeySequence__StandardKey(22)
	QKeySequence__FindNext                 QKeySequence__StandardKey = QKeySequence__StandardKey(23)
	QKeySequence__FindPrevious             QKeySequence__StandardKey = QKeySequence__StandardKey(24)
	QKeySequence__Replace                  QKeySequence__StandardKey = QKeySequence__StandardKey(25)
	QKeySequence__SelectAll                QKeySequence__StandardKey = QKeySequence__StandardKey(26)
	QKeySequence__Bold                     QKeySequence__StandardKey = QKeySequence__StandardKey(27)
	QKeySequence__Italic                   QKeySequence__StandardKey = QKeySequence__StandardKey(28)
	QKeySequence__Underline                QKeySequence__StandardKey = QKeySequence__StandardKey(29)
	QKeySequence__MoveToNextChar           QKeySequence__StandardKey = QKeySequence__StandardKey(30)
	QKeySequence__MoveToPreviousChar       QKeySequence__StandardKey = QKeySequence__StandardKey(31)
	QKeySequence__MoveToNextWord           QKeySequence__StandardKey = QKeySequence__StandardKey(32)
	QKeySequence__MoveToPreviousWord       QKeySequence__StandardKey = QKeySequence__StandardKey(33)
	QKeySequence__MoveToNextLine           QKeySequence__StandardKey = QKeySequence__StandardKey(34)
	QKeySequence__MoveToPreviousLine       QKeySequence__StandardKey = QKeySequence__StandardKey(35)
	QKeySequence__MoveToNextPage           QKeySequence__StandardKey = QKeySequence__StandardKey(36)
	QKeySequence__MoveToPreviousPage       QKeySequence__StandardKey = QKeySequence__StandardKey(37)
	QKeySequence__MoveToStartOfLine        QKeySequence__StandardKey = QKeySequence__StandardKey(38)
	QKeySequence__MoveToEndOfLine          QKeySequence__StandardKey = QKeySequence__StandardKey(39)
	QKeySequence__MoveToStartOfBlock       QKeySequence__StandardKey = QKeySequence__StandardKey(40)
	QKeySequence__MoveToEndOfBlock         QKeySequence__StandardKey = QKeySequence__StandardKey(41)
	QKeySequence__MoveToStartOfDocument    QKeySequence__StandardKey = QKeySequence__StandardKey(42)
	QKeySequence__MoveToEndOfDocument      QKeySequence__StandardKey = QKeySequence__StandardKey(43)
	QKeySequence__SelectNextChar           QKeySequence__StandardKey = QKeySequence__StandardKey(44)
	QKeySequence__SelectPreviousChar       QKeySequence__StandardKey = QKeySequence__StandardKey(45)
	QKeySequence__SelectNextWord           QKeySequence__StandardKey = QKeySequence__StandardKey(46)
	QKeySequence__SelectPreviousWord       QKeySequence__StandardKey = QKeySequence__StandardKey(47)
	QKeySequence__SelectNextLine           QKeySequence__StandardKey = QKeySequence__StandardKey(48)
	QKeySequence__SelectPreviousLine       QKeySequence__StandardKey = QKeySequence__StandardKey(49)
	QKeySequence__SelectNextPage           QKeySequence__StandardKey = QKeySequence__StandardKey(50)
	QKeySequence__SelectPreviousPage       QKeySequence__StandardKey = QKeySequence__StandardKey(51)
	QKeySequence__SelectStartOfLine        QKeySequence__StandardKey = QKeySequence__StandardKey(52)
	QKeySequence__SelectEndOfLine          QKeySequence__StandardKey = QKeySequence__StandardKey(53)
	QKeySequence__SelectStartOfBlock       QKeySequence__StandardKey = QKeySequence__StandardKey(54)
	QKeySequence__SelectEndOfBlock         QKeySequence__StandardKey = QKeySequence__StandardKey(55)
	QKeySequence__SelectStartOfDocument    QKeySequence__StandardKey = QKeySequence__StandardKey(56)
	QKeySequence__SelectEndOfDocument      QKeySequence__StandardKey = QKeySequence__StandardKey(57)
	QKeySequence__DeleteStartOfWord        QKeySequence__StandardKey = QKeySequence__StandardKey(58)
	QKeySequence__DeleteEndOfWord          QKeySequence__StandardKey = QKeySequence__StandardKey(59)
	QKeySequence__DeleteEndOfLine          QKeySequence__StandardKey = QKeySequence__StandardKey(60)
	QKeySequence__InsertParagraphSeparator QKeySequence__StandardKey = QKeySequence__StandardKey(61)
	QKeySequence__InsertLineSeparator      QKeySequence__StandardKey = QKeySequence__StandardKey(62)
	QKeySequence__SaveAs                   QKeySequence__StandardKey = QKeySequence__StandardKey(63)
	QKeySequence__Preferences              QKeySequence__StandardKey = QKeySequence__StandardKey(64)
	QKeySequence__Quit                     QKeySequence__StandardKey = QKeySequence__StandardKey(65)
	QKeySequence__FullScreen               QKeySequence__StandardKey = QKeySequence__StandardKey(66)
	QKeySequence__Deselect                 QKeySequence__StandardKey = QKeySequence__StandardKey(67)
	QKeySequence__DeleteCompleteLine       QKeySequence__StandardKey = QKeySequence__StandardKey(68)
	QKeySequence__Backspace                QKeySequence__StandardKey = QKeySequence__StandardKey(69)
	QKeySequence__Cancel                   QKeySequence__StandardKey = QKeySequence__StandardKey(70)
)

func NewQKeySequence() *QKeySequence {
	tmpValue := NewQKeySequenceFromPointer(C.QKeySequence_NewQKeySequence())
	qt.SetFinalizer(tmpValue, (*QKeySequence).DestroyQKeySequence)
	return tmpValue
}

func NewQKeySequence5(key QKeySequence__StandardKey) *QKeySequence {
	tmpValue := NewQKeySequenceFromPointer(C.QKeySequence_NewQKeySequence5(C.longlong(key)))
	qt.SetFinalizer(tmpValue, (*QKeySequence).DestroyQKeySequence)
	return tmpValue
}

func NewQKeySequence4(keysequence QKeySequence_ITF) *QKeySequence {
	tmpValue := NewQKeySequenceFromPointer(C.QKeySequence_NewQKeySequence4(PointerFromQKeySequence(keysequence)))
	qt.SetFinalizer(tmpValue, (*QKeySequence).DestroyQKeySequence)
	return tmpValue
}

func NewQKeySequence2(key string, format QKeySequence__SequenceFormat) *QKeySequence {
	var keyC *C.char
	if key != "" {
		keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
	}
	tmpValue := NewQKeySequenceFromPointer(C.QKeySequence_NewQKeySequence2(C.struct_QtGui_PackedString{data: keyC, len: C.longlong(len(key))}, C.longlong(format)))
	qt.SetFinalizer(tmpValue, (*QKeySequence).DestroyQKeySequence)
	return tmpValue
}

func NewQKeySequence3(k1 int, k2 int, k3 int, k4 int) *QKeySequence {
	tmpValue := NewQKeySequenceFromPointer(C.QKeySequence_NewQKeySequence3(C.int(int32(k1)), C.int(int32(k2)), C.int(int32(k3)), C.int(int32(k4))))
	qt.SetFinalizer(tmpValue, (*QKeySequence).DestroyQKeySequence)
	return tmpValue
}

func (ptr *QKeySequence) DestroyQKeySequence() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QKeySequence_DestroyQKeySequence(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QKeySequence) ToString(format QKeySequence__SequenceFormat) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QKeySequence_ToString(ptr.Pointer(), C.longlong(format)))
	}
	return ""
}

func (ptr *QKeySequence) __keyBindings_atList(i int) *QKeySequence {
	if ptr.Pointer() != nil {
		tmpValue := NewQKeySequenceFromPointer(C.QKeySequence___keyBindings_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QKeySequence).DestroyQKeySequence)
		return tmpValue
	}
	return nil
}

func (ptr *QKeySequence) __keyBindings_setList(i QKeySequence_ITF) {
	if ptr.Pointer() != nil {
		C.QKeySequence___keyBindings_setList(ptr.Pointer(), PointerFromQKeySequence(i))
	}
}

func (ptr *QKeySequence) __keyBindings_newList() unsafe.Pointer {
	return C.QKeySequence___keyBindings_newList(ptr.Pointer())
}

func (ptr *QKeySequence) __listFromString_atList(i int) *QKeySequence {
	if ptr.Pointer() != nil {
		tmpValue := NewQKeySequenceFromPointer(C.QKeySequence___listFromString_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QKeySequence).DestroyQKeySequence)
		return tmpValue
	}
	return nil
}

func (ptr *QKeySequence) __listFromString_setList(i QKeySequence_ITF) {
	if ptr.Pointer() != nil {
		C.QKeySequence___listFromString_setList(ptr.Pointer(), PointerFromQKeySequence(i))
	}
}

func (ptr *QKeySequence) __listFromString_newList() unsafe.Pointer {
	return C.QKeySequence___listFromString_newList(ptr.Pointer())
}

func (ptr *QKeySequence) __listToString_list_atList(i int) *QKeySequence {
	if ptr.Pointer() != nil {
		tmpValue := NewQKeySequenceFromPointer(C.QKeySequence___listToString_list_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QKeySequence).DestroyQKeySequence)
		return tmpValue
	}
	return nil
}

func (ptr *QKeySequence) __listToString_list_setList(i QKeySequence_ITF) {
	if ptr.Pointer() != nil {
		C.QKeySequence___listToString_list_setList(ptr.Pointer(), PointerFromQKeySequence(i))
	}
}

func (ptr *QKeySequence) __listToString_list_newList() unsafe.Pointer {
	return C.QKeySequence___listToString_list_newList(ptr.Pointer())
}

type QMatrix struct {
	ptr unsafe.Pointer
}

type QMatrix_ITF interface {
	QMatrix_PTR() *QMatrix
}

func (ptr *QMatrix) QMatrix_PTR() *QMatrix {
	return ptr
}

func (ptr *QMatrix) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QMatrix) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQMatrix(ptr QMatrix_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMatrix_PTR().Pointer()
	}
	return nil
}

func NewQMatrixFromPointer(ptr unsafe.Pointer) (n *QMatrix) {
	n = new(QMatrix)
	n.SetPointer(ptr)
	return
}
func (ptr *QMatrix) DestroyQMatrix() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func (ptr *QMatrix) Translate(dx float64, dy float64) *QMatrix {
	if ptr.Pointer() != nil {
		tmpValue := NewQMatrixFromPointer(C.QMatrix_Translate(ptr.Pointer(), C.double(dx), C.double(dy)))
		qt.SetFinalizer(tmpValue, (*QMatrix).DestroyQMatrix)
		return tmpValue
	}
	return nil
}

func NewQMatrix() *QMatrix {
	tmpValue := NewQMatrixFromPointer(C.QMatrix_NewQMatrix())
	qt.SetFinalizer(tmpValue, (*QMatrix).DestroyQMatrix)
	return tmpValue
}

func NewQMatrix4(other QMatrix_ITF) *QMatrix {
	tmpValue := NewQMatrixFromPointer(C.QMatrix_NewQMatrix4(PointerFromQMatrix(other)))
	qt.SetFinalizer(tmpValue, (*QMatrix).DestroyQMatrix)
	return tmpValue
}

func NewQMatrix5(matrix QMatrix_ITF) *QMatrix {
	tmpValue := NewQMatrixFromPointer(C.QMatrix_NewQMatrix5(PointerFromQMatrix(matrix)))
	qt.SetFinalizer(tmpValue, (*QMatrix).DestroyQMatrix)
	return tmpValue
}

func NewQMatrix3(m11 float64, m12 float64, m21 float64, m22 float64, dx float64, dy float64) *QMatrix {
	tmpValue := NewQMatrixFromPointer(C.QMatrix_NewQMatrix3(C.double(m11), C.double(m12), C.double(m21), C.double(m22), C.double(dx), C.double(dy)))
	qt.SetFinalizer(tmpValue, (*QMatrix).DestroyQMatrix)
	return tmpValue
}

func (ptr *QMatrix) Reset() {
	if ptr.Pointer() != nil {
		C.QMatrix_Reset(ptr.Pointer())
	}
}

func (ptr *QMatrix) Map5(line core.QLine_ITF) *core.QLine {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQLineFromPointer(C.QMatrix_Map5(ptr.Pointer(), core.PointerFromQLine(line)))
		qt.SetFinalizer(tmpValue, (*core.QLine).DestroyQLine)
		return tmpValue
	}
	return nil
}

func (ptr *QMatrix) Map6(line core.QLineF_ITF) *core.QLineF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQLineFFromPointer(C.QMatrix_Map6(ptr.Pointer(), core.PointerFromQLineF(line)))
		qt.SetFinalizer(tmpValue, (*core.QLineF).DestroyQLineF)
		return tmpValue
	}
	return nil
}

func (ptr *QMatrix) Map10(path QPainterPath_ITF) *QPainterPath {
	if ptr.Pointer() != nil {
		tmpValue := NewQPainterPathFromPointer(C.QMatrix_Map10(ptr.Pointer(), PointerFromQPainterPath(path)))
		qt.SetFinalizer(tmpValue, (*QPainterPath).DestroyQPainterPath)
		return tmpValue
	}
	return nil
}

func (ptr *QMatrix) Map3(point core.QPoint_ITF) *core.QPoint {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFromPointer(C.QMatrix_Map3(ptr.Pointer(), core.PointerFromQPoint(point)))
		qt.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QMatrix) Map4(point core.QPointF_ITF) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QMatrix_Map4(ptr.Pointer(), core.PointerFromQPointF(point)))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QMatrix) Map8(polygon QPolygon_ITF) *QPolygon {
	if ptr.Pointer() != nil {
		tmpValue := NewQPolygonFromPointer(C.QMatrix_Map8(ptr.Pointer(), PointerFromQPolygon(polygon)))
		qt.SetFinalizer(tmpValue, (*QPolygon).DestroyQPolygon)
		return tmpValue
	}
	return nil
}

func (ptr *QMatrix) Map7(polygon QPolygonF_ITF) *QPolygonF {
	if ptr.Pointer() != nil {
		tmpValue := NewQPolygonFFromPointer(C.QMatrix_Map7(ptr.Pointer(), PointerFromQPolygonF(polygon)))
		qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
		return tmpValue
	}
	return nil
}

func (ptr *QMatrix) Map9(region QRegion_ITF) *QRegion {
	if ptr.Pointer() != nil {
		tmpValue := NewQRegionFromPointer(C.QMatrix_Map9(ptr.Pointer(), PointerFromQRegion(region)))
		qt.SetFinalizer(tmpValue, (*QRegion).DestroyQRegion)
		return tmpValue
	}
	return nil
}

func (ptr *QMatrix) Map2(x int, y int, tx int, ty int) {
	if ptr.Pointer() != nil {
		C.QMatrix_Map2(ptr.Pointer(), C.int(int32(x)), C.int(int32(y)), C.int(int32(tx)), C.int(int32(ty)))
	}
}

func (ptr *QMatrix) Map(x float64, y float64, tx float64, ty float64) {
	if ptr.Pointer() != nil {
		C.QMatrix_Map(ptr.Pointer(), C.double(x), C.double(y), C.double(tx), C.double(ty))
	}
}

type QMatrix4x4 struct {
	ptr unsafe.Pointer
}

type QMatrix4x4_ITF interface {
	QMatrix4x4_PTR() *QMatrix4x4
}

func (ptr *QMatrix4x4) QMatrix4x4_PTR() *QMatrix4x4 {
	return ptr
}

func (ptr *QMatrix4x4) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QMatrix4x4) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQMatrix4x4(ptr QMatrix4x4_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMatrix4x4_PTR().Pointer()
	}
	return nil
}

func NewQMatrix4x4FromPointer(ptr unsafe.Pointer) (n *QMatrix4x4) {
	n = new(QMatrix4x4)
	n.SetPointer(ptr)
	return
}
func (ptr *QMatrix4x4) DestroyQMatrix4x4() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQMatrix4x4() *QMatrix4x4 {
	tmpValue := NewQMatrix4x4FromPointer(C.QMatrix4x4_NewQMatrix4x4())
	qt.SetFinalizer(tmpValue, (*QMatrix4x4).DestroyQMatrix4x4)
	return tmpValue
}

func NewQMatrix4x48(matrix QMatrix_ITF) *QMatrix4x4 {
	tmpValue := NewQMatrix4x4FromPointer(C.QMatrix4x4_NewQMatrix4x48(PointerFromQMatrix(matrix)))
	qt.SetFinalizer(tmpValue, (*QMatrix4x4).DestroyQMatrix4x4)
	return tmpValue
}

func NewQMatrix4x47(transform QTransform_ITF) *QMatrix4x4 {
	tmpValue := NewQMatrix4x4FromPointer(C.QMatrix4x4_NewQMatrix4x47(PointerFromQTransform(transform)))
	qt.SetFinalizer(tmpValue, (*QMatrix4x4).DestroyQMatrix4x4)
	return tmpValue
}

func NewQMatrix4x43(values float32) *QMatrix4x4 {
	tmpValue := NewQMatrix4x4FromPointer(C.QMatrix4x4_NewQMatrix4x43(C.float(values)))
	qt.SetFinalizer(tmpValue, (*QMatrix4x4).DestroyQMatrix4x4)
	return tmpValue
}

func NewQMatrix4x44(m11 float32, m12 float32, m13 float32, m14 float32, m21 float32, m22 float32, m23 float32, m24 float32, m31 float32, m32 float32, m33 float32, m34 float32, m41 float32, m42 float32, m43 float32, m44 float32) *QMatrix4x4 {
	tmpValue := NewQMatrix4x4FromPointer(C.QMatrix4x4_NewQMatrix4x44(C.float(m11), C.float(m12), C.float(m13), C.float(m14), C.float(m21), C.float(m22), C.float(m23), C.float(m24), C.float(m31), C.float(m32), C.float(m33), C.float(m34), C.float(m41), C.float(m42), C.float(m43), C.float(m44)))
	qt.SetFinalizer(tmpValue, (*QMatrix4x4).DestroyQMatrix4x4)
	return tmpValue
}

func (ptr *QMatrix4x4) Data() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QMatrix4x4_Data(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMatrix4x4) Translate(vector QVector3D_ITF) {
	if ptr.Pointer() != nil {
		C.QMatrix4x4_Translate(ptr.Pointer(), PointerFromQVector3D(vector))
	}
}

func (ptr *QMatrix4x4) Translate2(x float32, y float32) {
	if ptr.Pointer() != nil {
		C.QMatrix4x4_Translate2(ptr.Pointer(), C.float(x), C.float(y))
	}
}

func (ptr *QMatrix4x4) Translate3(x float32, y float32, z float32) {
	if ptr.Pointer() != nil {
		C.QMatrix4x4_Translate3(ptr.Pointer(), C.float(x), C.float(y), C.float(z))
	}
}

func (ptr *QMatrix4x4) Map(point core.QPoint_ITF) *core.QPoint {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFromPointer(C.QMatrix4x4_Map(ptr.Pointer(), core.PointerFromQPoint(point)))
		qt.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QMatrix4x4) Map2(point core.QPointF_ITF) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QMatrix4x4_Map2(ptr.Pointer(), core.PointerFromQPointF(point)))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QMatrix4x4) Map3(point QVector3D_ITF) *QVector3D {
	if ptr.Pointer() != nil {
		tmpValue := NewQVector3DFromPointer(C.QMatrix4x4_Map3(ptr.Pointer(), PointerFromQVector3D(point)))
		qt.SetFinalizer(tmpValue, (*QVector3D).DestroyQVector3D)
		return tmpValue
	}
	return nil
}

func (ptr *QMatrix4x4) Map4(point QVector4D_ITF) *QVector4D {
	if ptr.Pointer() != nil {
		tmpValue := NewQVector4DFromPointer(C.QMatrix4x4_Map4(ptr.Pointer(), PointerFromQVector4D(point)))
		qt.SetFinalizer(tmpValue, (*QVector4D).DestroyQVector4D)
		return tmpValue
	}
	return nil
}

func (ptr *QMatrix4x4) Data2() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QMatrix4x4_Data2(ptr.Pointer()))
	}
	return 0
}

//go:generate stringer -type=QMovie__CacheMode
//QMovie::CacheMode
type QMovie__CacheMode int64

const (
	QMovie__CacheNone QMovie__CacheMode = QMovie__CacheMode(0)
	QMovie__CacheAll  QMovie__CacheMode = QMovie__CacheMode(1)
)

//go:generate stringer -type=QMovie__MovieState
//QMovie::MovieState
type QMovie__MovieState int64

const (
	QMovie__NotRunning QMovie__MovieState = QMovie__MovieState(0)
	QMovie__Paused     QMovie__MovieState = QMovie__MovieState(1)
	QMovie__Running    QMovie__MovieState = QMovie__MovieState(2)
)

//go:generate stringer -type=QOpenGLBuffer__Access
//QOpenGLBuffer::Access
type QOpenGLBuffer__Access int64

const (
	QOpenGLBuffer__ReadOnly  QOpenGLBuffer__Access = QOpenGLBuffer__Access(0x88B8)
	QOpenGLBuffer__WriteOnly QOpenGLBuffer__Access = QOpenGLBuffer__Access(0x88B9)
	QOpenGLBuffer__ReadWrite QOpenGLBuffer__Access = QOpenGLBuffer__Access(0x88BA)
)

//go:generate stringer -type=QOpenGLBuffer__RangeAccessFlag
//QOpenGLBuffer::RangeAccessFlag
type QOpenGLBuffer__RangeAccessFlag int64

const (
	QOpenGLBuffer__RangeRead             QOpenGLBuffer__RangeAccessFlag = QOpenGLBuffer__RangeAccessFlag(0x0001)
	QOpenGLBuffer__RangeWrite            QOpenGLBuffer__RangeAccessFlag = QOpenGLBuffer__RangeAccessFlag(0x0002)
	QOpenGLBuffer__RangeInvalidate       QOpenGLBuffer__RangeAccessFlag = QOpenGLBuffer__RangeAccessFlag(0x0004)
	QOpenGLBuffer__RangeInvalidateBuffer QOpenGLBuffer__RangeAccessFlag = QOpenGLBuffer__RangeAccessFlag(0x0008)
	QOpenGLBuffer__RangeFlushExplicit    QOpenGLBuffer__RangeAccessFlag = QOpenGLBuffer__RangeAccessFlag(0x0010)
	QOpenGLBuffer__RangeUnsynchronized   QOpenGLBuffer__RangeAccessFlag = QOpenGLBuffer__RangeAccessFlag(0x0020)
)

//go:generate stringer -type=QOpenGLBuffer__Type
//QOpenGLBuffer::Type
type QOpenGLBuffer__Type int64

const (
	QOpenGLBuffer__VertexBuffer      QOpenGLBuffer__Type = QOpenGLBuffer__Type(0x8892)
	QOpenGLBuffer__IndexBuffer       QOpenGLBuffer__Type = QOpenGLBuffer__Type(0x8893)
	QOpenGLBuffer__PixelPackBuffer   QOpenGLBuffer__Type = QOpenGLBuffer__Type(0x88EB)
	QOpenGLBuffer__PixelUnpackBuffer QOpenGLBuffer__Type = QOpenGLBuffer__Type(0x88EC)
)

//go:generate stringer -type=QOpenGLBuffer__UsagePattern
//QOpenGLBuffer::UsagePattern
type QOpenGLBuffer__UsagePattern int64

const (
	QOpenGLBuffer__StreamDraw  QOpenGLBuffer__UsagePattern = QOpenGLBuffer__UsagePattern(0x88E0)
	QOpenGLBuffer__StreamRead  QOpenGLBuffer__UsagePattern = QOpenGLBuffer__UsagePattern(0x88E1)
	QOpenGLBuffer__StreamCopy  QOpenGLBuffer__UsagePattern = QOpenGLBuffer__UsagePattern(0x88E2)
	QOpenGLBuffer__StaticDraw  QOpenGLBuffer__UsagePattern = QOpenGLBuffer__UsagePattern(0x88E4)
	QOpenGLBuffer__StaticRead  QOpenGLBuffer__UsagePattern = QOpenGLBuffer__UsagePattern(0x88E5)
	QOpenGLBuffer__StaticCopy  QOpenGLBuffer__UsagePattern = QOpenGLBuffer__UsagePattern(0x88E6)
	QOpenGLBuffer__DynamicDraw QOpenGLBuffer__UsagePattern = QOpenGLBuffer__UsagePattern(0x88E8)
	QOpenGLBuffer__DynamicRead QOpenGLBuffer__UsagePattern = QOpenGLBuffer__UsagePattern(0x88E9)
	QOpenGLBuffer__DynamicCopy QOpenGLBuffer__UsagePattern = QOpenGLBuffer__UsagePattern(0x88EA)
)

//go:generate stringer -type=QOpenGLContext__OpenGLModuleType
//QOpenGLContext::OpenGLModuleType
type QOpenGLContext__OpenGLModuleType int64

const (
	QOpenGLContext__LibGL   QOpenGLContext__OpenGLModuleType = QOpenGLContext__OpenGLModuleType(0)
	QOpenGLContext__LibGLES QOpenGLContext__OpenGLModuleType = QOpenGLContext__OpenGLModuleType(1)
)

//go:generate stringer -type=QOpenGLDebugLogger__LoggingMode
//QOpenGLDebugLogger::LoggingMode
type QOpenGLDebugLogger__LoggingMode int64

const (
	QOpenGLDebugLogger__AsynchronousLogging QOpenGLDebugLogger__LoggingMode = QOpenGLDebugLogger__LoggingMode(0)
	QOpenGLDebugLogger__SynchronousLogging  QOpenGLDebugLogger__LoggingMode = QOpenGLDebugLogger__LoggingMode(1)
)

//go:generate stringer -type=QOpenGLDebugMessage__Severity
//QOpenGLDebugMessage::Severity
type QOpenGLDebugMessage__Severity int64

const (
	QOpenGLDebugMessage__InvalidSeverity      QOpenGLDebugMessage__Severity = QOpenGLDebugMessage__Severity(0x00000000)
	QOpenGLDebugMessage__HighSeverity         QOpenGLDebugMessage__Severity = QOpenGLDebugMessage__Severity(0x00000001)
	QOpenGLDebugMessage__MediumSeverity       QOpenGLDebugMessage__Severity = QOpenGLDebugMessage__Severity(0x00000002)
	QOpenGLDebugMessage__LowSeverity          QOpenGLDebugMessage__Severity = QOpenGLDebugMessage__Severity(0x00000004)
	QOpenGLDebugMessage__NotificationSeverity QOpenGLDebugMessage__Severity = QOpenGLDebugMessage__Severity(0x00000008)
	QOpenGLDebugMessage__LastSeverity         QOpenGLDebugMessage__Severity = QOpenGLDebugMessage__Severity(QOpenGLDebugMessage__NotificationSeverity)
	QOpenGLDebugMessage__AnySeverity          QOpenGLDebugMessage__Severity = QOpenGLDebugMessage__Severity(0xffffffff)
)

//go:generate stringer -type=QOpenGLDebugMessage__Source
//QOpenGLDebugMessage::Source
type QOpenGLDebugMessage__Source int64

const (
	QOpenGLDebugMessage__InvalidSource        QOpenGLDebugMessage__Source = QOpenGLDebugMessage__Source(0x00000000)
	QOpenGLDebugMessage__APISource            QOpenGLDebugMessage__Source = QOpenGLDebugMessage__Source(0x00000001)
	QOpenGLDebugMessage__WindowSystemSource   QOpenGLDebugMessage__Source = QOpenGLDebugMessage__Source(0x00000002)
	QOpenGLDebugMessage__ShaderCompilerSource QOpenGLDebugMessage__Source = QOpenGLDebugMessage__Source(0x00000004)
	QOpenGLDebugMessage__ThirdPartySource     QOpenGLDebugMessage__Source = QOpenGLDebugMessage__Source(0x00000008)
	QOpenGLDebugMessage__ApplicationSource    QOpenGLDebugMessage__Source = QOpenGLDebugMessage__Source(0x00000010)
	QOpenGLDebugMessage__OtherSource          QOpenGLDebugMessage__Source = QOpenGLDebugMessage__Source(0x00000020)
	QOpenGLDebugMessage__LastSource           QOpenGLDebugMessage__Source = QOpenGLDebugMessage__Source(QOpenGLDebugMessage__OtherSource)
	QOpenGLDebugMessage__AnySource            QOpenGLDebugMessage__Source = QOpenGLDebugMessage__Source(0xffffffff)
)

//go:generate stringer -type=QOpenGLDebugMessage__Type
//QOpenGLDebugMessage::Type
type QOpenGLDebugMessage__Type int64

const (
	QOpenGLDebugMessage__InvalidType            QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(0x00000000)
	QOpenGLDebugMessage__ErrorType              QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(0x00000001)
	QOpenGLDebugMessage__DeprecatedBehaviorType QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(0x00000002)
	QOpenGLDebugMessage__UndefinedBehaviorType  QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(0x00000004)
	QOpenGLDebugMessage__PortabilityType        QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(0x00000008)
	QOpenGLDebugMessage__PerformanceType        QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(0x00000010)
	QOpenGLDebugMessage__OtherType              QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(0x00000020)
	QOpenGLDebugMessage__MarkerType             QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(0x00000040)
	QOpenGLDebugMessage__GroupPushType          QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(0x00000080)
	QOpenGLDebugMessage__GroupPopType           QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(0x00000100)
	QOpenGLDebugMessage__LastType               QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(QOpenGLDebugMessage__GroupPopType)
	QOpenGLDebugMessage__AnyType                QOpenGLDebugMessage__Type = QOpenGLDebugMessage__Type(0xffffffff)
)

//go:generate stringer -type=QOpenGLFramebufferObject__Attachment
//QOpenGLFramebufferObject::Attachment
type QOpenGLFramebufferObject__Attachment int64

const (
	QOpenGLFramebufferObject__NoAttachment         QOpenGLFramebufferObject__Attachment = QOpenGLFramebufferObject__Attachment(0)
	QOpenGLFramebufferObject__CombinedDepthStencil QOpenGLFramebufferObject__Attachment = QOpenGLFramebufferObject__Attachment(1)
	QOpenGLFramebufferObject__Depth                QOpenGLFramebufferObject__Attachment = QOpenGLFramebufferObject__Attachment(2)
)

//go:generate stringer -type=QOpenGLFramebufferObject__FramebufferRestorePolicy
//QOpenGLFramebufferObject::FramebufferRestorePolicy
type QOpenGLFramebufferObject__FramebufferRestorePolicy int64

const (
	QOpenGLFramebufferObject__DontRestoreFramebufferBinding      QOpenGLFramebufferObject__FramebufferRestorePolicy = QOpenGLFramebufferObject__FramebufferRestorePolicy(0)
	QOpenGLFramebufferObject__RestoreFramebufferBindingToDefault QOpenGLFramebufferObject__FramebufferRestorePolicy = QOpenGLFramebufferObject__FramebufferRestorePolicy(1)
	QOpenGLFramebufferObject__RestoreFrameBufferBinding          QOpenGLFramebufferObject__FramebufferRestorePolicy = QOpenGLFramebufferObject__FramebufferRestorePolicy(2)
)

//go:generate stringer -type=QOpenGLFunctions__OpenGLFeature
//QOpenGLFunctions::OpenGLFeature
type QOpenGLFunctions__OpenGLFeature int64

const (
	QOpenGLFunctions__Multitexture          QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0001)
	QOpenGLFunctions__Shaders               QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0002)
	QOpenGLFunctions__Buffers               QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0004)
	QOpenGLFunctions__Framebuffers          QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0008)
	QOpenGLFunctions__BlendColor            QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0010)
	QOpenGLFunctions__BlendEquation         QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0020)
	QOpenGLFunctions__BlendEquationSeparate QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0040)
	QOpenGLFunctions__BlendFuncSeparate     QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0080)
	QOpenGLFunctions__BlendSubtract         QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0100)
	QOpenGLFunctions__CompressedTextures    QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0200)
	QOpenGLFunctions__Multisample           QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0400)
	QOpenGLFunctions__StencilSeparate       QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x0800)
	QOpenGLFunctions__NPOTTextures          QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x1000)
	QOpenGLFunctions__NPOTTextureRepeat     QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x2000)
	QOpenGLFunctions__FixedFunctionPipeline QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x4000)
	QOpenGLFunctions__TextureRGFormats      QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x8000)
	QOpenGLFunctions__MultipleRenderTargets QOpenGLFunctions__OpenGLFeature = QOpenGLFunctions__OpenGLFeature(0x10000)
)

//go:generate stringer -type=QOpenGLShader__ShaderTypeBit
//QOpenGLShader::ShaderTypeBit
type QOpenGLShader__ShaderTypeBit int64

const (
	QOpenGLShader__Vertex                 QOpenGLShader__ShaderTypeBit = QOpenGLShader__ShaderTypeBit(0x0001)
	QOpenGLShader__Fragment               QOpenGLShader__ShaderTypeBit = QOpenGLShader__ShaderTypeBit(0x0002)
	QOpenGLShader__Geometry               QOpenGLShader__ShaderTypeBit = QOpenGLShader__ShaderTypeBit(0x0004)
	QOpenGLShader__TessellationControl    QOpenGLShader__ShaderTypeBit = QOpenGLShader__ShaderTypeBit(0x0008)
	QOpenGLShader__TessellationEvaluation QOpenGLShader__ShaderTypeBit = QOpenGLShader__ShaderTypeBit(0x0010)
	QOpenGLShader__Compute                QOpenGLShader__ShaderTypeBit = QOpenGLShader__ShaderTypeBit(0x0020)
)

//go:generate stringer -type=QOpenGLTexture__BindingTarget
//QOpenGLTexture::BindingTarget
type QOpenGLTexture__BindingTarget int64

const (
	QOpenGLTexture__BindingTarget1D                 QOpenGLTexture__BindingTarget = QOpenGLTexture__BindingTarget(0x8068)
	QOpenGLTexture__BindingTarget1DArray            QOpenGLTexture__BindingTarget = QOpenGLTexture__BindingTarget(0x8C1C)
	QOpenGLTexture__BindingTarget2D                 QOpenGLTexture__BindingTarget = QOpenGLTexture__BindingTarget(0x8069)
	QOpenGLTexture__BindingTarget2DArray            QOpenGLTexture__BindingTarget = QOpenGLTexture__BindingTarget(0x8C1D)
	QOpenGLTexture__BindingTarget3D                 QOpenGLTexture__BindingTarget = QOpenGLTexture__BindingTarget(0x806A)
	QOpenGLTexture__BindingTargetCubeMap            QOpenGLTexture__BindingTarget = QOpenGLTexture__BindingTarget(0x8514)
	QOpenGLTexture__BindingTargetCubeMapArray       QOpenGLTexture__BindingTarget = QOpenGLTexture__BindingTarget(0x900A)
	QOpenGLTexture__BindingTarget2DMultisample      QOpenGLTexture__BindingTarget = QOpenGLTexture__BindingTarget(0x9104)
	QOpenGLTexture__BindingTarget2DMultisampleArray QOpenGLTexture__BindingTarget = QOpenGLTexture__BindingTarget(0x9105)
	QOpenGLTexture__BindingTargetRectangle          QOpenGLTexture__BindingTarget = QOpenGLTexture__BindingTarget(0x84F6)
	QOpenGLTexture__BindingTargetBuffer             QOpenGLTexture__BindingTarget = QOpenGLTexture__BindingTarget(0x8C2C)
)

//go:generate stringer -type=QOpenGLTexture__ComparisonFunction
//QOpenGLTexture::ComparisonFunction
type QOpenGLTexture__ComparisonFunction int64

const (
	QOpenGLTexture__CompareLessEqual    QOpenGLTexture__ComparisonFunction = QOpenGLTexture__ComparisonFunction(0x0203)
	QOpenGLTexture__CompareGreaterEqual QOpenGLTexture__ComparisonFunction = QOpenGLTexture__ComparisonFunction(0x0206)
	QOpenGLTexture__CompareLess         QOpenGLTexture__ComparisonFunction = QOpenGLTexture__ComparisonFunction(0x0201)
	QOpenGLTexture__CompareGreater      QOpenGLTexture__ComparisonFunction = QOpenGLTexture__ComparisonFunction(0x0204)
	QOpenGLTexture__CompareEqual        QOpenGLTexture__ComparisonFunction = QOpenGLTexture__ComparisonFunction(0x0202)
	QOpenGLTexture__CommpareNotEqual    QOpenGLTexture__ComparisonFunction = QOpenGLTexture__ComparisonFunction(0x0205)
	QOpenGLTexture__CompareAlways       QOpenGLTexture__ComparisonFunction = QOpenGLTexture__ComparisonFunction(0x0207)
	QOpenGLTexture__CompareNever        QOpenGLTexture__ComparisonFunction = QOpenGLTexture__ComparisonFunction(0x0200)
)

//go:generate stringer -type=QOpenGLTexture__ComparisonMode
//QOpenGLTexture::ComparisonMode
type QOpenGLTexture__ComparisonMode int64

const (
	QOpenGLTexture__CompareRefToTexture QOpenGLTexture__ComparisonMode = QOpenGLTexture__ComparisonMode(0x884E)
	QOpenGLTexture__CompareNone         QOpenGLTexture__ComparisonMode = QOpenGLTexture__ComparisonMode(0x0000)
)

//go:generate stringer -type=QOpenGLTexture__CoordinateDirection
//QOpenGLTexture::CoordinateDirection
type QOpenGLTexture__CoordinateDirection int64

const (
	QOpenGLTexture__DirectionS QOpenGLTexture__CoordinateDirection = QOpenGLTexture__CoordinateDirection(0x2802)
	QOpenGLTexture__DirectionT QOpenGLTexture__CoordinateDirection = QOpenGLTexture__CoordinateDirection(0x2803)
	QOpenGLTexture__DirectionR QOpenGLTexture__CoordinateDirection = QOpenGLTexture__CoordinateDirection(0x8072)
)

//go:generate stringer -type=QOpenGLTexture__CubeMapFace
//QOpenGLTexture::CubeMapFace
type QOpenGLTexture__CubeMapFace int64

const (
	QOpenGLTexture__CubeMapPositiveX QOpenGLTexture__CubeMapFace = QOpenGLTexture__CubeMapFace(0x8515)
	QOpenGLTexture__CubeMapNegativeX QOpenGLTexture__CubeMapFace = QOpenGLTexture__CubeMapFace(0x8516)
	QOpenGLTexture__CubeMapPositiveY QOpenGLTexture__CubeMapFace = QOpenGLTexture__CubeMapFace(0x8517)
	QOpenGLTexture__CubeMapNegativeY QOpenGLTexture__CubeMapFace = QOpenGLTexture__CubeMapFace(0x8518)
	QOpenGLTexture__CubeMapPositiveZ QOpenGLTexture__CubeMapFace = QOpenGLTexture__CubeMapFace(0x8519)
	QOpenGLTexture__CubeMapNegativeZ QOpenGLTexture__CubeMapFace = QOpenGLTexture__CubeMapFace(0x851A)
)

//go:generate stringer -type=QOpenGLTexture__DepthStencilMode
//QOpenGLTexture::DepthStencilMode
type QOpenGLTexture__DepthStencilMode int64

const (
	QOpenGLTexture__DepthMode   QOpenGLTexture__DepthStencilMode = QOpenGLTexture__DepthStencilMode(0x1902)
	QOpenGLTexture__StencilMode QOpenGLTexture__DepthStencilMode = QOpenGLTexture__DepthStencilMode(0x1901)
)

//go:generate stringer -type=QOpenGLTexture__Feature
//QOpenGLTexture::Feature
type QOpenGLTexture__Feature int64

const (
	QOpenGLTexture__ImmutableStorage            QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000001)
	QOpenGLTexture__ImmutableMultisampleStorage QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000002)
	QOpenGLTexture__TextureRectangle            QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000004)
	QOpenGLTexture__TextureArrays               QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000008)
	QOpenGLTexture__Texture3D                   QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000010)
	QOpenGLTexture__TextureMultisample          QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000020)
	QOpenGLTexture__TextureBuffer               QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000040)
	QOpenGLTexture__TextureCubeMapArrays        QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000080)
	QOpenGLTexture__Swizzle                     QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000100)
	QOpenGLTexture__StencilTexturing            QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000200)
	QOpenGLTexture__AnisotropicFiltering        QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000400)
	QOpenGLTexture__NPOTTextures                QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00000800)
	QOpenGLTexture__NPOTTextureRepeat           QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00001000)
	QOpenGLTexture__Texture1D                   QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00002000)
	QOpenGLTexture__TextureComparisonOperators  QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00004000)
	QOpenGLTexture__TextureMipMapLevel          QOpenGLTexture__Feature = QOpenGLTexture__Feature(0x00008000)
)

//go:generate stringer -type=QOpenGLTexture__Filter
//QOpenGLTexture::Filter
type QOpenGLTexture__Filter int64

const (
	QOpenGLTexture__Nearest              QOpenGLTexture__Filter = QOpenGLTexture__Filter(0x2600)
	QOpenGLTexture__Linear               QOpenGLTexture__Filter = QOpenGLTexture__Filter(0x2601)
	QOpenGLTexture__NearestMipMapNearest QOpenGLTexture__Filter = QOpenGLTexture__Filter(0x2700)
	QOpenGLTexture__NearestMipMapLinear  QOpenGLTexture__Filter = QOpenGLTexture__Filter(0x2702)
	QOpenGLTexture__LinearMipMapNearest  QOpenGLTexture__Filter = QOpenGLTexture__Filter(0x2701)
	QOpenGLTexture__LinearMipMapLinear   QOpenGLTexture__Filter = QOpenGLTexture__Filter(0x2703)
)

//go:generate stringer -type=QOpenGLTexture__MipMapGeneration
//QOpenGLTexture::MipMapGeneration
type QOpenGLTexture__MipMapGeneration int64

const (
	QOpenGLTexture__GenerateMipMaps     QOpenGLTexture__MipMapGeneration = QOpenGLTexture__MipMapGeneration(0)
	QOpenGLTexture__DontGenerateMipMaps QOpenGLTexture__MipMapGeneration = QOpenGLTexture__MipMapGeneration(1)
)

//go:generate stringer -type=QOpenGLTexture__PixelFormat
//QOpenGLTexture::PixelFormat
type QOpenGLTexture__PixelFormat int64

const (
	QOpenGLTexture__NoSourceFormat QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0)
	QOpenGLTexture__Red            QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x1903)
	QOpenGLTexture__RG             QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x8227)
	QOpenGLTexture__RGB            QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x1907)
	QOpenGLTexture__BGR            QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x80E0)
	QOpenGLTexture__RGBA           QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x1908)
	QOpenGLTexture__BGRA           QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x80E1)
	QOpenGLTexture__Red_Integer    QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x8D94)
	QOpenGLTexture__RG_Integer     QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x8228)
	QOpenGLTexture__RGB_Integer    QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x8D98)
	QOpenGLTexture__BGR_Integer    QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x8D9A)
	QOpenGLTexture__RGBA_Integer   QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x8D99)
	QOpenGLTexture__BGRA_Integer   QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x8D9B)
	QOpenGLTexture__Stencil        QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x1901)
	QOpenGLTexture__Depth          QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x1902)
	QOpenGLTexture__DepthStencil   QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x84F9)
	QOpenGLTexture__Alpha          QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x1906)
	QOpenGLTexture__Luminance      QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x1909)
	QOpenGLTexture__LuminanceAlpha QOpenGLTexture__PixelFormat = QOpenGLTexture__PixelFormat(0x190A)
)

//go:generate stringer -type=QOpenGLTexture__PixelType
//QOpenGLTexture::PixelType
type QOpenGLTexture__PixelType int64

const (
	QOpenGLTexture__NoPixelType               QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0)
	QOpenGLTexture__Int8                      QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x1400)
	QOpenGLTexture__UInt8                     QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x1401)
	QOpenGLTexture__Int16                     QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x1402)
	QOpenGLTexture__UInt16                    QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x1403)
	QOpenGLTexture__Int32                     QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x1404)
	QOpenGLTexture__UInt32                    QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x1405)
	QOpenGLTexture__Float16                   QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x140B)
	QOpenGLTexture__Float16OES                QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8D61)
	QOpenGLTexture__Float32                   QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x1406)
	QOpenGLTexture__UInt32_RGB9_E5            QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8C3E)
	QOpenGLTexture__UInt32_RG11B10F           QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8C3B)
	QOpenGLTexture__UInt8_RG3B2               QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8032)
	QOpenGLTexture__UInt8_RG3B2_Rev           QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8362)
	QOpenGLTexture__UInt16_RGB5A1             QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8034)
	QOpenGLTexture__UInt16_RGB5A1_Rev         QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8366)
	QOpenGLTexture__UInt16_R5G6B5             QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8363)
	QOpenGLTexture__UInt16_R5G6B5_Rev         QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8364)
	QOpenGLTexture__UInt16_RGBA4              QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8033)
	QOpenGLTexture__UInt16_RGBA4_Rev          QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8365)
	QOpenGLTexture__UInt32_RGBA8              QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8035)
	QOpenGLTexture__UInt32_RGBA8_Rev          QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8367)
	QOpenGLTexture__UInt32_RGB10A2            QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8036)
	QOpenGLTexture__UInt32_RGB10A2_Rev        QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8368)
	QOpenGLTexture__UInt32_D24S8              QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x84FA)
	QOpenGLTexture__Float32_D32_UInt32_S8_X24 QOpenGLTexture__PixelType = QOpenGLTexture__PixelType(0x8DAD)
)

//go:generate stringer -type=QOpenGLTexture__SwizzleComponent
//QOpenGLTexture::SwizzleComponent
type QOpenGLTexture__SwizzleComponent int64

const (
	QOpenGLTexture__SwizzleRed   QOpenGLTexture__SwizzleComponent = QOpenGLTexture__SwizzleComponent(0x8E42)
	QOpenGLTexture__SwizzleGreen QOpenGLTexture__SwizzleComponent = QOpenGLTexture__SwizzleComponent(0x8E43)
	QOpenGLTexture__SwizzleBlue  QOpenGLTexture__SwizzleComponent = QOpenGLTexture__SwizzleComponent(0x8E44)
	QOpenGLTexture__SwizzleAlpha QOpenGLTexture__SwizzleComponent = QOpenGLTexture__SwizzleComponent(0x8E45)
)

//go:generate stringer -type=QOpenGLTexture__SwizzleValue
//QOpenGLTexture::SwizzleValue
type QOpenGLTexture__SwizzleValue int64

const (
	QOpenGLTexture__RedValue   QOpenGLTexture__SwizzleValue = QOpenGLTexture__SwizzleValue(0x1903)
	QOpenGLTexture__GreenValue QOpenGLTexture__SwizzleValue = QOpenGLTexture__SwizzleValue(0x1904)
	QOpenGLTexture__BlueValue  QOpenGLTexture__SwizzleValue = QOpenGLTexture__SwizzleValue(0x1905)
	QOpenGLTexture__AlphaValue QOpenGLTexture__SwizzleValue = QOpenGLTexture__SwizzleValue(0x1906)
	QOpenGLTexture__ZeroValue  QOpenGLTexture__SwizzleValue = QOpenGLTexture__SwizzleValue(0)
	QOpenGLTexture__OneValue   QOpenGLTexture__SwizzleValue = QOpenGLTexture__SwizzleValue(1)
)

//go:generate stringer -type=QOpenGLTexture__Target
//QOpenGLTexture::Target
type QOpenGLTexture__Target int64

const (
	QOpenGLTexture__Target1D                 QOpenGLTexture__Target = QOpenGLTexture__Target(0x0DE0)
	QOpenGLTexture__Target1DArray            QOpenGLTexture__Target = QOpenGLTexture__Target(0x8C18)
	QOpenGLTexture__Target2D                 QOpenGLTexture__Target = QOpenGLTexture__Target(0x0DE1)
	QOpenGLTexture__Target2DArray            QOpenGLTexture__Target = QOpenGLTexture__Target(0x8C1A)
	QOpenGLTexture__Target3D                 QOpenGLTexture__Target = QOpenGLTexture__Target(0x806F)
	QOpenGLTexture__TargetCubeMap            QOpenGLTexture__Target = QOpenGLTexture__Target(0x8513)
	QOpenGLTexture__TargetCubeMapArray       QOpenGLTexture__Target = QOpenGLTexture__Target(0x9009)
	QOpenGLTexture__Target2DMultisample      QOpenGLTexture__Target = QOpenGLTexture__Target(0x9100)
	QOpenGLTexture__Target2DMultisampleArray QOpenGLTexture__Target = QOpenGLTexture__Target(0x9102)
	QOpenGLTexture__TargetRectangle          QOpenGLTexture__Target = QOpenGLTexture__Target(0x84F5)
	QOpenGLTexture__TargetBuffer             QOpenGLTexture__Target = QOpenGLTexture__Target(0x8C2A)
)

//go:generate stringer -type=QOpenGLTexture__TextureFormat
//QOpenGLTexture::TextureFormat
type QOpenGLTexture__TextureFormat int64

const (
	QOpenGLTexture__NoFormat                       QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0)
	QOpenGLTexture__R8_UNorm                       QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8229)
	QOpenGLTexture__RG8_UNorm                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x822B)
	QOpenGLTexture__RGB8_UNorm                     QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8051)
	QOpenGLTexture__RGBA8_UNorm                    QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8058)
	QOpenGLTexture__R16_UNorm                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x822A)
	QOpenGLTexture__RG16_UNorm                     QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x822C)
	QOpenGLTexture__RGB16_UNorm                    QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8054)
	QOpenGLTexture__RGBA16_UNorm                   QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x805B)
	QOpenGLTexture__R8_SNorm                       QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8F94)
	QOpenGLTexture__RG8_SNorm                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8F95)
	QOpenGLTexture__RGB8_SNorm                     QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8F96)
	QOpenGLTexture__RGBA8_SNorm                    QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8F97)
	QOpenGLTexture__R16_SNorm                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8F98)
	QOpenGLTexture__RG16_SNorm                     QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8F99)
	QOpenGLTexture__RGB16_SNorm                    QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8F9A)
	QOpenGLTexture__RGBA16_SNorm                   QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8F9B)
	QOpenGLTexture__R8U                            QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8232)
	QOpenGLTexture__RG8U                           QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8238)
	QOpenGLTexture__RGB8U                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D7D)
	QOpenGLTexture__RGBA8U                         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D7C)
	QOpenGLTexture__R16U                           QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8234)
	QOpenGLTexture__RG16U                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x823A)
	QOpenGLTexture__RGB16U                         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D77)
	QOpenGLTexture__RGBA16U                        QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D76)
	QOpenGLTexture__R32U                           QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8236)
	QOpenGLTexture__RG32U                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x823C)
	QOpenGLTexture__RGB32U                         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D71)
	QOpenGLTexture__RGBA32U                        QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D70)
	QOpenGLTexture__R8I                            QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8231)
	QOpenGLTexture__RG8I                           QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8237)
	QOpenGLTexture__RGB8I                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D8F)
	QOpenGLTexture__RGBA8I                         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D8E)
	QOpenGLTexture__R16I                           QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8233)
	QOpenGLTexture__RG16I                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8239)
	QOpenGLTexture__RGB16I                         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D89)
	QOpenGLTexture__RGBA16I                        QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D88)
	QOpenGLTexture__R32I                           QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8235)
	QOpenGLTexture__RG32I                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x823B)
	QOpenGLTexture__RGB32I                         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D83)
	QOpenGLTexture__RGBA32I                        QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D82)
	QOpenGLTexture__R16F                           QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x822D)
	QOpenGLTexture__RG16F                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x822F)
	QOpenGLTexture__RGB16F                         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x881B)
	QOpenGLTexture__RGBA16F                        QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x881A)
	QOpenGLTexture__R32F                           QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x822E)
	QOpenGLTexture__RG32F                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8230)
	QOpenGLTexture__RGB32F                         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8815)
	QOpenGLTexture__RGBA32F                        QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8814)
	QOpenGLTexture__RGB9E5                         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8C3D)
	QOpenGLTexture__RG11B10F                       QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8C3A)
	QOpenGLTexture__RG3B2                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x2A10)
	QOpenGLTexture__R5G6B5                         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D62)
	QOpenGLTexture__RGB5A1                         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8057)
	QOpenGLTexture__RGBA4                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8056)
	QOpenGLTexture__RGB10A2                        QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x906F)
	QOpenGLTexture__D16                            QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x81A5)
	QOpenGLTexture__D24                            QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x81A6)
	QOpenGLTexture__D24S8                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x88F0)
	QOpenGLTexture__D32                            QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x81A7)
	QOpenGLTexture__D32F                           QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8CAC)
	QOpenGLTexture__D32FS8X24                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8CAD)
	QOpenGLTexture__S8                             QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D48)
	QOpenGLTexture__RGB_DXT1                       QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x83F0)
	QOpenGLTexture__RGBA_DXT1                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x83F1)
	QOpenGLTexture__RGBA_DXT3                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x83F2)
	QOpenGLTexture__RGBA_DXT5                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x83F3)
	QOpenGLTexture__R_ATI1N_UNorm                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8DBB)
	QOpenGLTexture__R_ATI1N_SNorm                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8DBC)
	QOpenGLTexture__RG_ATI2N_UNorm                 QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8DBD)
	QOpenGLTexture__RG_ATI2N_SNorm                 QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8DBE)
	QOpenGLTexture__RGB_BP_UNSIGNED_FLOAT          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8E8F)
	QOpenGLTexture__RGB_BP_SIGNED_FLOAT            QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8E8E)
	QOpenGLTexture__RGB_BP_UNorm                   QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8E8C)
	QOpenGLTexture__R11_EAC_UNorm                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x9270)
	QOpenGLTexture__R11_EAC_SNorm                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x9271)
	QOpenGLTexture__RG11_EAC_UNorm                 QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x9272)
	QOpenGLTexture__RG11_EAC_SNorm                 QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x9273)
	QOpenGLTexture__RGB8_ETC2                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x9274)
	QOpenGLTexture__SRGB8_ETC2                     QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x9275)
	QOpenGLTexture__RGB8_PunchThrough_Alpha1_ETC2  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x9276)
	QOpenGLTexture__SRGB8_PunchThrough_Alpha1_ETC2 QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x9277)
	QOpenGLTexture__RGBA8_ETC2_EAC                 QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x9278)
	QOpenGLTexture__SRGB8_Alpha8_ETC2_EAC          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x9279)
	QOpenGLTexture__RGB8_ETC1                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8D64)
	QOpenGLTexture__RGBA_ASTC_4x4                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93B0)
	QOpenGLTexture__RGBA_ASTC_5x4                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93B1)
	QOpenGLTexture__RGBA_ASTC_5x5                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93B2)
	QOpenGLTexture__RGBA_ASTC_6x5                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93B3)
	QOpenGLTexture__RGBA_ASTC_6x6                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93B4)
	QOpenGLTexture__RGBA_ASTC_8x5                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93B5)
	QOpenGLTexture__RGBA_ASTC_8x6                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93B6)
	QOpenGLTexture__RGBA_ASTC_8x8                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93B7)
	QOpenGLTexture__RGBA_ASTC_10x5                 QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93B8)
	QOpenGLTexture__RGBA_ASTC_10x6                 QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93B9)
	QOpenGLTexture__RGBA_ASTC_10x8                 QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93BA)
	QOpenGLTexture__RGBA_ASTC_10x10                QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93BB)
	QOpenGLTexture__RGBA_ASTC_12x10                QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93BC)
	QOpenGLTexture__RGBA_ASTC_12x12                QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93BD)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_4x4          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93D0)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_5x4          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93D1)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_5x5          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93D2)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_6x5          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93D3)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_6x6          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93D4)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_8x5          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93D5)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_8x6          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93D6)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_8x8          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93D7)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_10x5         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93D8)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_10x6         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93D9)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_10x8         QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93DA)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_10x10        QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93DB)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_12x10        QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93DC)
	QOpenGLTexture__SRGB8_Alpha8_ASTC_12x12        QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x93DD)
	QOpenGLTexture__SRGB8                          QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8C41)
	QOpenGLTexture__SRGB8_Alpha8                   QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8C43)
	QOpenGLTexture__SRGB_DXT1                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8C4C)
	QOpenGLTexture__SRGB_Alpha_DXT1                QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8C4D)
	QOpenGLTexture__SRGB_Alpha_DXT3                QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8C4E)
	QOpenGLTexture__SRGB_Alpha_DXT5                QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8C4F)
	QOpenGLTexture__SRGB_BP_UNorm                  QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x8E8D)
	QOpenGLTexture__DepthFormat                    QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x1902)
	QOpenGLTexture__AlphaFormat                    QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x1906)
	QOpenGLTexture__RGBFormat                      QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x1907)
	QOpenGLTexture__RGBAFormat                     QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x1908)
	QOpenGLTexture__LuminanceFormat                QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x1909)
	QOpenGLTexture__LuminanceAlphaFormat           QOpenGLTexture__TextureFormat = QOpenGLTexture__TextureFormat(0x190A)
)

//go:generate stringer -type=QOpenGLTexture__TextureUnitReset
//QOpenGLTexture::TextureUnitReset
type QOpenGLTexture__TextureUnitReset int64

const (
	QOpenGLTexture__ResetTextureUnit     QOpenGLTexture__TextureUnitReset = QOpenGLTexture__TextureUnitReset(0)
	QOpenGLTexture__DontResetTextureUnit QOpenGLTexture__TextureUnitReset = QOpenGLTexture__TextureUnitReset(1)
)

//go:generate stringer -type=QOpenGLTexture__WrapMode
//QOpenGLTexture::WrapMode
type QOpenGLTexture__WrapMode int64

const (
	QOpenGLTexture__Repeat         QOpenGLTexture__WrapMode = QOpenGLTexture__WrapMode(0x2901)
	QOpenGLTexture__MirroredRepeat QOpenGLTexture__WrapMode = QOpenGLTexture__WrapMode(0x8370)
	QOpenGLTexture__ClampToEdge    QOpenGLTexture__WrapMode = QOpenGLTexture__WrapMode(0x812F)
	QOpenGLTexture__ClampToBorder  QOpenGLTexture__WrapMode = QOpenGLTexture__WrapMode(0x812D)
)

//go:generate stringer -type=QOpenGLTextureBlitter__Origin
//QOpenGLTextureBlitter::Origin
type QOpenGLTextureBlitter__Origin int64

const (
	QOpenGLTextureBlitter__OriginBottomLeft QOpenGLTextureBlitter__Origin = QOpenGLTextureBlitter__Origin(0)
	QOpenGLTextureBlitter__OriginTopLeft    QOpenGLTextureBlitter__Origin = QOpenGLTextureBlitter__Origin(1)
)

//go:generate stringer -type=QOpenGLWindow__UpdateBehavior
//QOpenGLWindow::UpdateBehavior
type QOpenGLWindow__UpdateBehavior int64

const (
	QOpenGLWindow__NoPartialUpdate    QOpenGLWindow__UpdateBehavior = QOpenGLWindow__UpdateBehavior(0)
	QOpenGLWindow__PartialUpdateBlit  QOpenGLWindow__UpdateBehavior = QOpenGLWindow__UpdateBehavior(1)
	QOpenGLWindow__PartialUpdateBlend QOpenGLWindow__UpdateBehavior = QOpenGLWindow__UpdateBehavior(2)
)

//go:generate stringer -type=QPageLayout__Mode
//QPageLayout::Mode
type QPageLayout__Mode int64

const (
	QPageLayout__StandardMode QPageLayout__Mode = QPageLayout__Mode(0)
	QPageLayout__FullPageMode QPageLayout__Mode = QPageLayout__Mode(1)
)

//go:generate stringer -type=QPageLayout__Orientation
//QPageLayout::Orientation
type QPageLayout__Orientation int64

const (
	QPageLayout__Portrait  QPageLayout__Orientation = QPageLayout__Orientation(0)
	QPageLayout__Landscape QPageLayout__Orientation = QPageLayout__Orientation(1)
)

//go:generate stringer -type=QPageLayout__Unit
//QPageLayout::Unit
type QPageLayout__Unit int64

const (
	QPageLayout__Millimeter QPageLayout__Unit = QPageLayout__Unit(0)
	QPageLayout__Point      QPageLayout__Unit = QPageLayout__Unit(1)
	QPageLayout__Inch       QPageLayout__Unit = QPageLayout__Unit(2)
	QPageLayout__Pica       QPageLayout__Unit = QPageLayout__Unit(3)
	QPageLayout__Didot      QPageLayout__Unit = QPageLayout__Unit(4)
	QPageLayout__Cicero     QPageLayout__Unit = QPageLayout__Unit(5)
)

//go:generate stringer -type=QPageSize__PageSizeId
//QPageSize::PageSizeId
type QPageSize__PageSizeId int64

const (
	QPageSize__A4                 QPageSize__PageSizeId = QPageSize__PageSizeId(0)
	QPageSize__B5                 QPageSize__PageSizeId = QPageSize__PageSizeId(1)
	QPageSize__Letter             QPageSize__PageSizeId = QPageSize__PageSizeId(2)
	QPageSize__Legal              QPageSize__PageSizeId = QPageSize__PageSizeId(3)
	QPageSize__Executive          QPageSize__PageSizeId = QPageSize__PageSizeId(4)
	QPageSize__A0                 QPageSize__PageSizeId = QPageSize__PageSizeId(5)
	QPageSize__A1                 QPageSize__PageSizeId = QPageSize__PageSizeId(6)
	QPageSize__A2                 QPageSize__PageSizeId = QPageSize__PageSizeId(7)
	QPageSize__A3                 QPageSize__PageSizeId = QPageSize__PageSizeId(8)
	QPageSize__A5                 QPageSize__PageSizeId = QPageSize__PageSizeId(9)
	QPageSize__A6                 QPageSize__PageSizeId = QPageSize__PageSizeId(10)
	QPageSize__A7                 QPageSize__PageSizeId = QPageSize__PageSizeId(11)
	QPageSize__A8                 QPageSize__PageSizeId = QPageSize__PageSizeId(12)
	QPageSize__A9                 QPageSize__PageSizeId = QPageSize__PageSizeId(13)
	QPageSize__B0                 QPageSize__PageSizeId = QPageSize__PageSizeId(14)
	QPageSize__B1                 QPageSize__PageSizeId = QPageSize__PageSizeId(15)
	QPageSize__B10                QPageSize__PageSizeId = QPageSize__PageSizeId(16)
	QPageSize__B2                 QPageSize__PageSizeId = QPageSize__PageSizeId(17)
	QPageSize__B3                 QPageSize__PageSizeId = QPageSize__PageSizeId(18)
	QPageSize__B4                 QPageSize__PageSizeId = QPageSize__PageSizeId(19)
	QPageSize__B6                 QPageSize__PageSizeId = QPageSize__PageSizeId(20)
	QPageSize__B7                 QPageSize__PageSizeId = QPageSize__PageSizeId(21)
	QPageSize__B8                 QPageSize__PageSizeId = QPageSize__PageSizeId(22)
	QPageSize__B9                 QPageSize__PageSizeId = QPageSize__PageSizeId(23)
	QPageSize__C5E                QPageSize__PageSizeId = QPageSize__PageSizeId(24)
	QPageSize__Comm10E            QPageSize__PageSizeId = QPageSize__PageSizeId(25)
	QPageSize__DLE                QPageSize__PageSizeId = QPageSize__PageSizeId(26)
	QPageSize__Folio              QPageSize__PageSizeId = QPageSize__PageSizeId(27)
	QPageSize__Ledger             QPageSize__PageSizeId = QPageSize__PageSizeId(28)
	QPageSize__Tabloid            QPageSize__PageSizeId = QPageSize__PageSizeId(29)
	QPageSize__Custom             QPageSize__PageSizeId = QPageSize__PageSizeId(30)
	QPageSize__A10                QPageSize__PageSizeId = QPageSize__PageSizeId(31)
	QPageSize__A3Extra            QPageSize__PageSizeId = QPageSize__PageSizeId(32)
	QPageSize__A4Extra            QPageSize__PageSizeId = QPageSize__PageSizeId(33)
	QPageSize__A4Plus             QPageSize__PageSizeId = QPageSize__PageSizeId(34)
	QPageSize__A4Small            QPageSize__PageSizeId = QPageSize__PageSizeId(35)
	QPageSize__A5Extra            QPageSize__PageSizeId = QPageSize__PageSizeId(36)
	QPageSize__B5Extra            QPageSize__PageSizeId = QPageSize__PageSizeId(37)
	QPageSize__JisB0              QPageSize__PageSizeId = QPageSize__PageSizeId(38)
	QPageSize__JisB1              QPageSize__PageSizeId = QPageSize__PageSizeId(39)
	QPageSize__JisB2              QPageSize__PageSizeId = QPageSize__PageSizeId(40)
	QPageSize__JisB3              QPageSize__PageSizeId = QPageSize__PageSizeId(41)
	QPageSize__JisB4              QPageSize__PageSizeId = QPageSize__PageSizeId(42)
	QPageSize__JisB5              QPageSize__PageSizeId = QPageSize__PageSizeId(43)
	QPageSize__JisB6              QPageSize__PageSizeId = QPageSize__PageSizeId(44)
	QPageSize__JisB7              QPageSize__PageSizeId = QPageSize__PageSizeId(45)
	QPageSize__JisB8              QPageSize__PageSizeId = QPageSize__PageSizeId(46)
	QPageSize__JisB9              QPageSize__PageSizeId = QPageSize__PageSizeId(47)
	QPageSize__JisB10             QPageSize__PageSizeId = QPageSize__PageSizeId(48)
	QPageSize__AnsiC              QPageSize__PageSizeId = QPageSize__PageSizeId(49)
	QPageSize__AnsiD              QPageSize__PageSizeId = QPageSize__PageSizeId(50)
	QPageSize__AnsiE              QPageSize__PageSizeId = QPageSize__PageSizeId(51)
	QPageSize__LegalExtra         QPageSize__PageSizeId = QPageSize__PageSizeId(52)
	QPageSize__LetterExtra        QPageSize__PageSizeId = QPageSize__PageSizeId(53)
	QPageSize__LetterPlus         QPageSize__PageSizeId = QPageSize__PageSizeId(54)
	QPageSize__LetterSmall        QPageSize__PageSizeId = QPageSize__PageSizeId(55)
	QPageSize__TabloidExtra       QPageSize__PageSizeId = QPageSize__PageSizeId(56)
	QPageSize__ArchA              QPageSize__PageSizeId = QPageSize__PageSizeId(57)
	QPageSize__ArchB              QPageSize__PageSizeId = QPageSize__PageSizeId(58)
	QPageSize__ArchC              QPageSize__PageSizeId = QPageSize__PageSizeId(59)
	QPageSize__ArchD              QPageSize__PageSizeId = QPageSize__PageSizeId(60)
	QPageSize__ArchE              QPageSize__PageSizeId = QPageSize__PageSizeId(61)
	QPageSize__Imperial7x9        QPageSize__PageSizeId = QPageSize__PageSizeId(62)
	QPageSize__Imperial8x10       QPageSize__PageSizeId = QPageSize__PageSizeId(63)
	QPageSize__Imperial9x11       QPageSize__PageSizeId = QPageSize__PageSizeId(64)
	QPageSize__Imperial9x12       QPageSize__PageSizeId = QPageSize__PageSizeId(65)
	QPageSize__Imperial10x11      QPageSize__PageSizeId = QPageSize__PageSizeId(66)
	QPageSize__Imperial10x13      QPageSize__PageSizeId = QPageSize__PageSizeId(67)
	QPageSize__Imperial10x14      QPageSize__PageSizeId = QPageSize__PageSizeId(68)
	QPageSize__Imperial12x11      QPageSize__PageSizeId = QPageSize__PageSizeId(69)
	QPageSize__Imperial15x11      QPageSize__PageSizeId = QPageSize__PageSizeId(70)
	QPageSize__ExecutiveStandard  QPageSize__PageSizeId = QPageSize__PageSizeId(71)
	QPageSize__Note               QPageSize__PageSizeId = QPageSize__PageSizeId(72)
	QPageSize__Quarto             QPageSize__PageSizeId = QPageSize__PageSizeId(73)
	QPageSize__Statement          QPageSize__PageSizeId = QPageSize__PageSizeId(74)
	QPageSize__SuperA             QPageSize__PageSizeId = QPageSize__PageSizeId(75)
	QPageSize__SuperB             QPageSize__PageSizeId = QPageSize__PageSizeId(76)
	QPageSize__Postcard           QPageSize__PageSizeId = QPageSize__PageSizeId(77)
	QPageSize__DoublePostcard     QPageSize__PageSizeId = QPageSize__PageSizeId(78)
	QPageSize__Prc16K             QPageSize__PageSizeId = QPageSize__PageSizeId(79)
	QPageSize__Prc32K             QPageSize__PageSizeId = QPageSize__PageSizeId(80)
	QPageSize__Prc32KBig          QPageSize__PageSizeId = QPageSize__PageSizeId(81)
	QPageSize__FanFoldUS          QPageSize__PageSizeId = QPageSize__PageSizeId(82)
	QPageSize__FanFoldGerman      QPageSize__PageSizeId = QPageSize__PageSizeId(83)
	QPageSize__FanFoldGermanLegal QPageSize__PageSizeId = QPageSize__PageSizeId(84)
	QPageSize__EnvelopeB4         QPageSize__PageSizeId = QPageSize__PageSizeId(85)
	QPageSize__EnvelopeB5         QPageSize__PageSizeId = QPageSize__PageSizeId(86)
	QPageSize__EnvelopeB6         QPageSize__PageSizeId = QPageSize__PageSizeId(87)
	QPageSize__EnvelopeC0         QPageSize__PageSizeId = QPageSize__PageSizeId(88)
	QPageSize__EnvelopeC1         QPageSize__PageSizeId = QPageSize__PageSizeId(89)
	QPageSize__EnvelopeC2         QPageSize__PageSizeId = QPageSize__PageSizeId(90)
	QPageSize__EnvelopeC3         QPageSize__PageSizeId = QPageSize__PageSizeId(91)
	QPageSize__EnvelopeC4         QPageSize__PageSizeId = QPageSize__PageSizeId(92)
	QPageSize__EnvelopeC6         QPageSize__PageSizeId = QPageSize__PageSizeId(93)
	QPageSize__EnvelopeC65        QPageSize__PageSizeId = QPageSize__PageSizeId(94)
	QPageSize__EnvelopeC7         QPageSize__PageSizeId = QPageSize__PageSizeId(95)
	QPageSize__Envelope9          QPageSize__PageSizeId = QPageSize__PageSizeId(96)
	QPageSize__Envelope11         QPageSize__PageSizeId = QPageSize__PageSizeId(97)
	QPageSize__Envelope12         QPageSize__PageSizeId = QPageSize__PageSizeId(98)
	QPageSize__Envelope14         QPageSize__PageSizeId = QPageSize__PageSizeId(99)
	QPageSize__EnvelopeMonarch    QPageSize__PageSizeId = QPageSize__PageSizeId(100)
	QPageSize__EnvelopePersonal   QPageSize__PageSizeId = QPageSize__PageSizeId(101)
	QPageSize__EnvelopeChou3      QPageSize__PageSizeId = QPageSize__PageSizeId(102)
	QPageSize__EnvelopeChou4      QPageSize__PageSizeId = QPageSize__PageSizeId(103)
	QPageSize__EnvelopeInvite     QPageSize__PageSizeId = QPageSize__PageSizeId(104)
	QPageSize__EnvelopeItalian    QPageSize__PageSizeId = QPageSize__PageSizeId(105)
	QPageSize__EnvelopeKaku2      QPageSize__PageSizeId = QPageSize__PageSizeId(106)
	QPageSize__EnvelopeKaku3      QPageSize__PageSizeId = QPageSize__PageSizeId(107)
	QPageSize__EnvelopePrc1       QPageSize__PageSizeId = QPageSize__PageSizeId(108)
	QPageSize__EnvelopePrc2       QPageSize__PageSizeId = QPageSize__PageSizeId(109)
	QPageSize__EnvelopePrc3       QPageSize__PageSizeId = QPageSize__PageSizeId(110)
	QPageSize__EnvelopePrc4       QPageSize__PageSizeId = QPageSize__PageSizeId(111)
	QPageSize__EnvelopePrc5       QPageSize__PageSizeId = QPageSize__PageSizeId(112)
	QPageSize__EnvelopePrc6       QPageSize__PageSizeId = QPageSize__PageSizeId(113)
	QPageSize__EnvelopePrc7       QPageSize__PageSizeId = QPageSize__PageSizeId(114)
	QPageSize__EnvelopePrc8       QPageSize__PageSizeId = QPageSize__PageSizeId(115)
	QPageSize__EnvelopePrc9       QPageSize__PageSizeId = QPageSize__PageSizeId(116)
	QPageSize__EnvelopePrc10      QPageSize__PageSizeId = QPageSize__PageSizeId(117)
	QPageSize__EnvelopeYou4       QPageSize__PageSizeId = QPageSize__PageSizeId(118)
	QPageSize__LastPageSize       QPageSize__PageSizeId = QPageSize__PageSizeId(QPageSize__EnvelopeYou4)
	QPageSize__NPageSize          QPageSize__PageSizeId = QPageSize__PageSizeId(QPageSize__LastPageSize)
	QPageSize__NPaperSize         QPageSize__PageSizeId = QPageSize__PageSizeId(QPageSize__LastPageSize)
	QPageSize__AnsiA              QPageSize__PageSizeId = QPageSize__PageSizeId(QPageSize__Letter)
	QPageSize__AnsiB              QPageSize__PageSizeId = QPageSize__PageSizeId(QPageSize__Ledger)
	QPageSize__EnvelopeC5         QPageSize__PageSizeId = QPageSize__PageSizeId(QPageSize__C5E)
	QPageSize__EnvelopeDL         QPageSize__PageSizeId = QPageSize__PageSizeId(QPageSize__DLE)
	QPageSize__Envelope10         QPageSize__PageSizeId = QPageSize__PageSizeId(QPageSize__Comm10E)
)

//go:generate stringer -type=QPageSize__SizeMatchPolicy
//QPageSize::SizeMatchPolicy
type QPageSize__SizeMatchPolicy int64

const (
	QPageSize__FuzzyMatch            QPageSize__SizeMatchPolicy = QPageSize__SizeMatchPolicy(0)
	QPageSize__FuzzyOrientationMatch QPageSize__SizeMatchPolicy = QPageSize__SizeMatchPolicy(1)
	QPageSize__ExactMatch            QPageSize__SizeMatchPolicy = QPageSize__SizeMatchPolicy(2)
)

//go:generate stringer -type=QPageSize__Unit
//QPageSize::Unit
type QPageSize__Unit int64

const (
	QPageSize__Millimeter QPageSize__Unit = QPageSize__Unit(0)
	QPageSize__Point      QPageSize__Unit = QPageSize__Unit(1)
	QPageSize__Inch       QPageSize__Unit = QPageSize__Unit(2)
	QPageSize__Pica       QPageSize__Unit = QPageSize__Unit(3)
	QPageSize__Didot      QPageSize__Unit = QPageSize__Unit(4)
	QPageSize__Cicero     QPageSize__Unit = QPageSize__Unit(5)
)

type QPagedPaintDevice struct {
	QPaintDevice
}

type QPagedPaintDevice_ITF interface {
	QPaintDevice_ITF
	QPagedPaintDevice_PTR() *QPagedPaintDevice
}

func (ptr *QPagedPaintDevice) QPagedPaintDevice_PTR() *QPagedPaintDevice {
	return ptr
}

func (ptr *QPagedPaintDevice) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintDevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QPagedPaintDevice) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPaintDevice_PTR().SetPointer(p)
	}
}

func PointerFromQPagedPaintDevice(ptr QPagedPaintDevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPagedPaintDevice_PTR().Pointer()
	}
	return nil
}

func NewQPagedPaintDeviceFromPointer(ptr unsafe.Pointer) (n *QPagedPaintDevice) {
	n = new(QPagedPaintDevice)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QPagedPaintDevice__PageSize
//QPagedPaintDevice::PageSize
type QPagedPaintDevice__PageSize int64

const (
	QPagedPaintDevice__A4                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(0)
	QPagedPaintDevice__B5                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(1)
	QPagedPaintDevice__Letter             QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(2)
	QPagedPaintDevice__Legal              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(3)
	QPagedPaintDevice__Executive          QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(4)
	QPagedPaintDevice__A0                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(5)
	QPagedPaintDevice__A1                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(6)
	QPagedPaintDevice__A2                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(7)
	QPagedPaintDevice__A3                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(8)
	QPagedPaintDevice__A5                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(9)
	QPagedPaintDevice__A6                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(10)
	QPagedPaintDevice__A7                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(11)
	QPagedPaintDevice__A8                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(12)
	QPagedPaintDevice__A9                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(13)
	QPagedPaintDevice__B0                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(14)
	QPagedPaintDevice__B1                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(15)
	QPagedPaintDevice__B10                QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(16)
	QPagedPaintDevice__B2                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(17)
	QPagedPaintDevice__B3                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(18)
	QPagedPaintDevice__B4                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(19)
	QPagedPaintDevice__B6                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(20)
	QPagedPaintDevice__B7                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(21)
	QPagedPaintDevice__B8                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(22)
	QPagedPaintDevice__B9                 QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(23)
	QPagedPaintDevice__C5E                QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(24)
	QPagedPaintDevice__Comm10E            QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(25)
	QPagedPaintDevice__DLE                QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(26)
	QPagedPaintDevice__Folio              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(27)
	QPagedPaintDevice__Ledger             QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(28)
	QPagedPaintDevice__Tabloid            QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(29)
	QPagedPaintDevice__Custom             QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(30)
	QPagedPaintDevice__A10                QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(31)
	QPagedPaintDevice__A3Extra            QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(32)
	QPagedPaintDevice__A4Extra            QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(33)
	QPagedPaintDevice__A4Plus             QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(34)
	QPagedPaintDevice__A4Small            QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(35)
	QPagedPaintDevice__A5Extra            QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(36)
	QPagedPaintDevice__B5Extra            QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(37)
	QPagedPaintDevice__JisB0              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(38)
	QPagedPaintDevice__JisB1              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(39)
	QPagedPaintDevice__JisB2              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(40)
	QPagedPaintDevice__JisB3              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(41)
	QPagedPaintDevice__JisB4              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(42)
	QPagedPaintDevice__JisB5              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(43)
	QPagedPaintDevice__JisB6              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(44)
	QPagedPaintDevice__JisB7              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(45)
	QPagedPaintDevice__JisB8              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(46)
	QPagedPaintDevice__JisB9              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(47)
	QPagedPaintDevice__JisB10             QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(48)
	QPagedPaintDevice__AnsiC              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(49)
	QPagedPaintDevice__AnsiD              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(50)
	QPagedPaintDevice__AnsiE              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(51)
	QPagedPaintDevice__LegalExtra         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(52)
	QPagedPaintDevice__LetterExtra        QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(53)
	QPagedPaintDevice__LetterPlus         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(54)
	QPagedPaintDevice__LetterSmall        QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(55)
	QPagedPaintDevice__TabloidExtra       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(56)
	QPagedPaintDevice__ArchA              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(57)
	QPagedPaintDevice__ArchB              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(58)
	QPagedPaintDevice__ArchC              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(59)
	QPagedPaintDevice__ArchD              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(60)
	QPagedPaintDevice__ArchE              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(61)
	QPagedPaintDevice__Imperial7x9        QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(62)
	QPagedPaintDevice__Imperial8x10       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(63)
	QPagedPaintDevice__Imperial9x11       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(64)
	QPagedPaintDevice__Imperial9x12       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(65)
	QPagedPaintDevice__Imperial10x11      QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(66)
	QPagedPaintDevice__Imperial10x13      QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(67)
	QPagedPaintDevice__Imperial10x14      QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(68)
	QPagedPaintDevice__Imperial12x11      QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(69)
	QPagedPaintDevice__Imperial15x11      QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(70)
	QPagedPaintDevice__ExecutiveStandard  QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(71)
	QPagedPaintDevice__Note               QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(72)
	QPagedPaintDevice__Quarto             QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(73)
	QPagedPaintDevice__Statement          QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(74)
	QPagedPaintDevice__SuperA             QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(75)
	QPagedPaintDevice__SuperB             QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(76)
	QPagedPaintDevice__Postcard           QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(77)
	QPagedPaintDevice__DoublePostcard     QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(78)
	QPagedPaintDevice__Prc16K             QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(79)
	QPagedPaintDevice__Prc32K             QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(80)
	QPagedPaintDevice__Prc32KBig          QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(81)
	QPagedPaintDevice__FanFoldUS          QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(82)
	QPagedPaintDevice__FanFoldGerman      QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(83)
	QPagedPaintDevice__FanFoldGermanLegal QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(84)
	QPagedPaintDevice__EnvelopeB4         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(85)
	QPagedPaintDevice__EnvelopeB5         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(86)
	QPagedPaintDevice__EnvelopeB6         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(87)
	QPagedPaintDevice__EnvelopeC0         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(88)
	QPagedPaintDevice__EnvelopeC1         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(89)
	QPagedPaintDevice__EnvelopeC2         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(90)
	QPagedPaintDevice__EnvelopeC3         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(91)
	QPagedPaintDevice__EnvelopeC4         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(92)
	QPagedPaintDevice__EnvelopeC6         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(93)
	QPagedPaintDevice__EnvelopeC65        QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(94)
	QPagedPaintDevice__EnvelopeC7         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(95)
	QPagedPaintDevice__Envelope9          QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(96)
	QPagedPaintDevice__Envelope11         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(97)
	QPagedPaintDevice__Envelope12         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(98)
	QPagedPaintDevice__Envelope14         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(99)
	QPagedPaintDevice__EnvelopeMonarch    QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(100)
	QPagedPaintDevice__EnvelopePersonal   QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(101)
	QPagedPaintDevice__EnvelopeChou3      QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(102)
	QPagedPaintDevice__EnvelopeChou4      QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(103)
	QPagedPaintDevice__EnvelopeInvite     QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(104)
	QPagedPaintDevice__EnvelopeItalian    QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(105)
	QPagedPaintDevice__EnvelopeKaku2      QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(106)
	QPagedPaintDevice__EnvelopeKaku3      QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(107)
	QPagedPaintDevice__EnvelopePrc1       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(108)
	QPagedPaintDevice__EnvelopePrc2       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(109)
	QPagedPaintDevice__EnvelopePrc3       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(110)
	QPagedPaintDevice__EnvelopePrc4       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(111)
	QPagedPaintDevice__EnvelopePrc5       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(112)
	QPagedPaintDevice__EnvelopePrc6       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(113)
	QPagedPaintDevice__EnvelopePrc7       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(114)
	QPagedPaintDevice__EnvelopePrc8       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(115)
	QPagedPaintDevice__EnvelopePrc9       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(116)
	QPagedPaintDevice__EnvelopePrc10      QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(117)
	QPagedPaintDevice__EnvelopeYou4       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(118)
	QPagedPaintDevice__LastPageSize       QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(QPagedPaintDevice__EnvelopeYou4)
	QPagedPaintDevice__NPageSize          QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(QPagedPaintDevice__LastPageSize)
	QPagedPaintDevice__NPaperSize         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(QPagedPaintDevice__LastPageSize)
	QPagedPaintDevice__AnsiA              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(QPagedPaintDevice__Letter)
	QPagedPaintDevice__AnsiB              QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(QPagedPaintDevice__Ledger)
	QPagedPaintDevice__EnvelopeC5         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(QPagedPaintDevice__C5E)
	QPagedPaintDevice__EnvelopeDL         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(QPagedPaintDevice__DLE)
	QPagedPaintDevice__Envelope10         QPagedPaintDevice__PageSize = QPagedPaintDevice__PageSize(QPagedPaintDevice__Comm10E)
)

func NewQPagedPaintDevice() *QPagedPaintDevice {
	return NewQPagedPaintDeviceFromPointer(C.QPagedPaintDevice_NewQPagedPaintDevice())
}

//export callbackQPagedPaintDevice_NewPage
func callbackQPagedPaintDevice_NewPage(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "newPage"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QPagedPaintDevice) ConnectNewPage(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "newPage"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "newPage", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "newPage", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPagedPaintDevice) DisconnectNewPage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "newPage")
	}
}

func (ptr *QPagedPaintDevice) NewPage() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPagedPaintDevice_NewPage(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPagedPaintDevice) DestroyQPagedPaintDevice() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPagedPaintDevice_DestroyQPagedPaintDevice(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQPagedPaintDevice_PaintEngine
func callbackQPagedPaintDevice_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return PointerFromQPaintEngine((*(*func() *QPaintEngine)(signal))())
	}

	return PointerFromQPaintEngine(NewQPagedPaintDeviceFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QPagedPaintDevice) PaintEngine() *QPaintEngine {
	if ptr.Pointer() != nil {
		return NewQPaintEngineFromPointer(C.QPagedPaintDevice_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPagedPaintDevice) PaintEngineDefault() *QPaintEngine {
	if ptr.Pointer() != nil {
		return NewQPaintEngineFromPointer(C.QPagedPaintDevice_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

type QPaintDevice struct {
	ptr unsafe.Pointer
}

type QPaintDevice_ITF interface {
	QPaintDevice_PTR() *QPaintDevice
}

func (ptr *QPaintDevice) QPaintDevice_PTR() *QPaintDevice {
	return ptr
}

func (ptr *QPaintDevice) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPaintDevice) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPaintDevice(ptr QPaintDevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintDevice_PTR().Pointer()
	}
	return nil
}

func NewQPaintDeviceFromPointer(ptr unsafe.Pointer) (n *QPaintDevice) {
	n = new(QPaintDevice)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QPaintDevice__PaintDeviceMetric
//QPaintDevice::PaintDeviceMetric
type QPaintDevice__PaintDeviceMetric int64

const (
	QPaintDevice__PdmWidth                  QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(1)
	QPaintDevice__PdmHeight                 QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(2)
	QPaintDevice__PdmWidthMM                QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(3)
	QPaintDevice__PdmHeightMM               QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(4)
	QPaintDevice__PdmNumColors              QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(5)
	QPaintDevice__PdmDepth                  QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(6)
	QPaintDevice__PdmDpiX                   QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(7)
	QPaintDevice__PdmDpiY                   QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(8)
	QPaintDevice__PdmPhysicalDpiX           QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(9)
	QPaintDevice__PdmPhysicalDpiY           QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(10)
	QPaintDevice__PdmDevicePixelRatio       QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(11)
	QPaintDevice__PdmDevicePixelRatioScaled QPaintDevice__PaintDeviceMetric = QPaintDevice__PaintDeviceMetric(12)
)

func NewQPaintDevice() *QPaintDevice {
	return NewQPaintDeviceFromPointer(C.QPaintDevice_NewQPaintDevice())
}

//export callbackQPaintDevice_DestroyQPaintDevice
func callbackQPaintDevice_DestroyQPaintDevice(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QPaintDevice"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQPaintDeviceFromPointer(ptr).DestroyQPaintDeviceDefault()
	}
}

func (ptr *QPaintDevice) ConnectDestroyQPaintDevice(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QPaintDevice"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QPaintDevice", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QPaintDevice", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPaintDevice) DisconnectDestroyQPaintDevice() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QPaintDevice")
	}
}

func (ptr *QPaintDevice) DestroyQPaintDevice() {
	if ptr.Pointer() != nil {
		C.QPaintDevice_DestroyQPaintDevice(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPaintDevice) DestroyQPaintDeviceDefault() {
	if ptr.Pointer() != nil {
		C.QPaintDevice_DestroyQPaintDeviceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQPaintDevice_PaintEngine
func callbackQPaintDevice_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return PointerFromQPaintEngine((*(*func() *QPaintEngine)(signal))())
	}

	return PointerFromQPaintEngine(NewQPaintEngine(0))
}

func (ptr *QPaintDevice) ConnectPaintEngine(f func() *QPaintEngine) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "paintEngine"); signal != nil {
			f := func() *QPaintEngine {
				(*(*func() *QPaintEngine)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "paintEngine", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "paintEngine", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPaintDevice) DisconnectPaintEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "paintEngine")
	}
}

func (ptr *QPaintDevice) PaintEngine() *QPaintEngine {
	if ptr.Pointer() != nil {
		return NewQPaintEngineFromPointer(C.QPaintDevice_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPaintDevice) Depth() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QPaintDevice_Depth(ptr.Pointer())))
	}
	return 0
}

type QPaintEngine struct {
	ptr unsafe.Pointer
}

type QPaintEngine_ITF interface {
	QPaintEngine_PTR() *QPaintEngine
}

func (ptr *QPaintEngine) QPaintEngine_PTR() *QPaintEngine {
	return ptr
}

func (ptr *QPaintEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPaintEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPaintEngine(ptr QPaintEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintEngine_PTR().Pointer()
	}
	return nil
}

func NewQPaintEngineFromPointer(ptr unsafe.Pointer) (n *QPaintEngine) {
	n = new(QPaintEngine)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QPaintEngine__DirtyFlag
//QPaintEngine::DirtyFlag
type QPaintEngine__DirtyFlag int64

const (
	QPaintEngine__DirtyPen             QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0001)
	QPaintEngine__DirtyBrush           QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0002)
	QPaintEngine__DirtyBrushOrigin     QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0004)
	QPaintEngine__DirtyFont            QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0008)
	QPaintEngine__DirtyBackground      QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0010)
	QPaintEngine__DirtyBackgroundMode  QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0020)
	QPaintEngine__DirtyTransform       QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0040)
	QPaintEngine__DirtyClipRegion      QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0080)
	QPaintEngine__DirtyClipPath        QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0100)
	QPaintEngine__DirtyHints           QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0200)
	QPaintEngine__DirtyCompositionMode QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0400)
	QPaintEngine__DirtyClipEnabled     QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x0800)
	QPaintEngine__DirtyOpacity         QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0x1000)
	QPaintEngine__AllDirty             QPaintEngine__DirtyFlag = QPaintEngine__DirtyFlag(0xffff)
)

//go:generate stringer -type=QPaintEngine__PaintEngineFeature
//QPaintEngine::PaintEngineFeature
type QPaintEngine__PaintEngineFeature int64

const (
	QPaintEngine__PrimitiveTransform          QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000001)
	QPaintEngine__PatternTransform            QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000002)
	QPaintEngine__PixmapTransform             QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000004)
	QPaintEngine__PatternBrush                QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000008)
	QPaintEngine__LinearGradientFill          QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000010)
	QPaintEngine__RadialGradientFill          QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000020)
	QPaintEngine__ConicalGradientFill         QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000040)
	QPaintEngine__AlphaBlend                  QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000080)
	QPaintEngine__PorterDuff                  QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000100)
	QPaintEngine__PainterPaths                QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000200)
	QPaintEngine__Antialiasing                QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000400)
	QPaintEngine__BrushStroke                 QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00000800)
	QPaintEngine__ConstantOpacity             QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00001000)
	QPaintEngine__MaskedBrush                 QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00002000)
	QPaintEngine__PerspectiveTransform        QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00004000)
	QPaintEngine__BlendModes                  QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00008000)
	QPaintEngine__ObjectBoundingModeGradients QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00010000)
	QPaintEngine__RasterOpModes               QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x00020000)
	QPaintEngine__PaintOutsidePaintEvent      QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0x20000000)
	QPaintEngine__AllFeatures                 QPaintEngine__PaintEngineFeature = QPaintEngine__PaintEngineFeature(0xffffffff)
)

//go:generate stringer -type=QPaintEngine__PolygonDrawMode
//QPaintEngine::PolygonDrawMode
type QPaintEngine__PolygonDrawMode int64

const (
	QPaintEngine__OddEvenMode  QPaintEngine__PolygonDrawMode = QPaintEngine__PolygonDrawMode(0)
	QPaintEngine__WindingMode  QPaintEngine__PolygonDrawMode = QPaintEngine__PolygonDrawMode(1)
	QPaintEngine__ConvexMode   QPaintEngine__PolygonDrawMode = QPaintEngine__PolygonDrawMode(2)
	QPaintEngine__PolylineMode QPaintEngine__PolygonDrawMode = QPaintEngine__PolygonDrawMode(3)
)

//go:generate stringer -type=QPaintEngine__Type
//QPaintEngine::Type
type QPaintEngine__Type int64

const (
	QPaintEngine__X11           QPaintEngine__Type = QPaintEngine__Type(0)
	QPaintEngine__Windows       QPaintEngine__Type = QPaintEngine__Type(1)
	QPaintEngine__QuickDraw     QPaintEngine__Type = QPaintEngine__Type(2)
	QPaintEngine__CoreGraphics  QPaintEngine__Type = QPaintEngine__Type(3)
	QPaintEngine__MacPrinter    QPaintEngine__Type = QPaintEngine__Type(4)
	QPaintEngine__QWindowSystem QPaintEngine__Type = QPaintEngine__Type(5)
	QPaintEngine__PostScript    QPaintEngine__Type = QPaintEngine__Type(6)
	QPaintEngine__OpenGL        QPaintEngine__Type = QPaintEngine__Type(7)
	QPaintEngine__Picture       QPaintEngine__Type = QPaintEngine__Type(8)
	QPaintEngine__SVG           QPaintEngine__Type = QPaintEngine__Type(9)
	QPaintEngine__Raster        QPaintEngine__Type = QPaintEngine__Type(10)
	QPaintEngine__Direct3D      QPaintEngine__Type = QPaintEngine__Type(11)
	QPaintEngine__Pdf           QPaintEngine__Type = QPaintEngine__Type(12)
	QPaintEngine__OpenVG        QPaintEngine__Type = QPaintEngine__Type(13)
	QPaintEngine__OpenGL2       QPaintEngine__Type = QPaintEngine__Type(14)
	QPaintEngine__PaintBuffer   QPaintEngine__Type = QPaintEngine__Type(15)
	QPaintEngine__Blitter       QPaintEngine__Type = QPaintEngine__Type(16)
	QPaintEngine__Direct2D      QPaintEngine__Type = QPaintEngine__Type(17)
	QPaintEngine__User          QPaintEngine__Type = QPaintEngine__Type(50)
	QPaintEngine__MaxUser       QPaintEngine__Type = QPaintEngine__Type(100)
)

//export callbackQPaintEngine_End
func callbackQPaintEngine_End(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "end"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QPaintEngine) ConnectEnd(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "end"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "end", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "end", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPaintEngine) DisconnectEnd() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "end")
	}
}

func (ptr *QPaintEngine) End() bool {
	if ptr.Pointer() != nil {
		return int8(C.QPaintEngine_End(ptr.Pointer())) != 0
	}
	return false
}

func NewQPaintEngine(caps QPaintEngine__PaintEngineFeature) *QPaintEngine {
	return NewQPaintEngineFromPointer(C.QPaintEngine_NewQPaintEngine(C.longlong(caps)))
}

//export callbackQPaintEngine_Begin
func callbackQPaintEngine_Begin(ptr unsafe.Pointer, pdev unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "begin"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*QPaintDevice) bool)(signal))(NewQPaintDeviceFromPointer(pdev)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QPaintEngine) ConnectBegin(f func(pdev *QPaintDevice) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "begin"); signal != nil {
			f := func(pdev *QPaintDevice) bool {
				(*(*func(*QPaintDevice) bool)(signal))(pdev)
				return f(pdev)
			}
			qt.ConnectSignal(ptr.Pointer(), "begin", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "begin", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPaintEngine) DisconnectBegin() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "begin")
	}
}

func (ptr *QPaintEngine) Begin(pdev QPaintDevice_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPaintEngine_Begin(ptr.Pointer(), PointerFromQPaintDevice(pdev))) != 0
	}
	return false
}

//export callbackQPaintEngine_DrawPixmap
func callbackQPaintEngine_DrawPixmap(ptr unsafe.Pointer, r unsafe.Pointer, pm unsafe.Pointer, sr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "drawPixmap"); signal != nil {
		(*(*func(*core.QRectF, *QPixmap, *core.QRectF))(signal))(core.NewQRectFFromPointer(r), NewQPixmapFromPointer(pm), core.NewQRectFFromPointer(sr))
	}

}

func (ptr *QPaintEngine) ConnectDrawPixmap(f func(r *core.QRectF, pm *QPixmap, sr *core.QRectF)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "drawPixmap"); signal != nil {
			f := func(r *core.QRectF, pm *QPixmap, sr *core.QRectF) {
				(*(*func(*core.QRectF, *QPixmap, *core.QRectF))(signal))(r, pm, sr)
				f(r, pm, sr)
			}
			qt.ConnectSignal(ptr.Pointer(), "drawPixmap", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "drawPixmap", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPaintEngine) DisconnectDrawPixmap() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "drawPixmap")
	}
}

func (ptr *QPaintEngine) DrawPixmap(r core.QRectF_ITF, pm QPixmap_ITF, sr core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawPixmap(ptr.Pointer(), core.PointerFromQRectF(r), PointerFromQPixmap(pm), core.PointerFromQRectF(sr))
	}
}

//export callbackQPaintEngine_UpdateState
func callbackQPaintEngine_UpdateState(ptr unsafe.Pointer, state unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updateState"); signal != nil {
		(*(*func(*QPaintEngineState))(signal))(NewQPaintEngineStateFromPointer(state))
	}

}

func (ptr *QPaintEngine) ConnectUpdateState(f func(state *QPaintEngineState)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "updateState"); signal != nil {
			f := func(state *QPaintEngineState) {
				(*(*func(*QPaintEngineState))(signal))(state)
				f(state)
			}
			qt.ConnectSignal(ptr.Pointer(), "updateState", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "updateState", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPaintEngine) DisconnectUpdateState() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "updateState")
	}
}

func (ptr *QPaintEngine) UpdateState(state QPaintEngineState_ITF) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_UpdateState(ptr.Pointer(), PointerFromQPaintEngineState(state))
	}
}

//export callbackQPaintEngine_DestroyQPaintEngine
func callbackQPaintEngine_DestroyQPaintEngine(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QPaintEngine"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQPaintEngineFromPointer(ptr).DestroyQPaintEngineDefault()
	}
}

func (ptr *QPaintEngine) ConnectDestroyQPaintEngine(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QPaintEngine"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QPaintEngine", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QPaintEngine", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPaintEngine) DisconnectDestroyQPaintEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QPaintEngine")
	}
}

func (ptr *QPaintEngine) DestroyQPaintEngine() {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DestroyQPaintEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPaintEngine) DestroyQPaintEngineDefault() {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DestroyQPaintEngineDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQPaintEngine_Type
func callbackQPaintEngine_Type(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "type"); signal != nil {
		return C.longlong((*(*func() QPaintEngine__Type)(signal))())
	}

	return C.longlong(0)
}

func (ptr *QPaintEngine) ConnectType(f func() QPaintEngine__Type) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "type"); signal != nil {
			f := func() QPaintEngine__Type {
				(*(*func() QPaintEngine__Type)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "type", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QPaintEngine) DisconnectType() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "type")
	}
}

func (ptr *QPaintEngine) Type() QPaintEngine__Type {
	if ptr.Pointer() != nil {
		return QPaintEngine__Type(C.QPaintEngine_Type(ptr.Pointer()))
	}
	return 0
}

type QPaintEngineState struct {
	ptr unsafe.Pointer
}

type QPaintEngineState_ITF interface {
	QPaintEngineState_PTR() *QPaintEngineState
}

func (ptr *QPaintEngineState) QPaintEngineState_PTR() *QPaintEngineState {
	return ptr
}

func (ptr *QPaintEngineState) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPaintEngineState) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPaintEngineState(ptr QPaintEngineState_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintEngineState_PTR().Pointer()
	}
	return nil
}

func NewQPaintEngineStateFromPointer(ptr unsafe.Pointer) (n *QPaintEngineState) {
	n = new(QPaintEngineState)
	n.SetPointer(ptr)
	return
}
func (ptr *QPaintEngineState) DestroyQPaintEngineState() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QPaintEvent struct {
	core.QEvent
}

type QPaintEvent_ITF interface {
	core.QEvent_ITF
	QPaintEvent_PTR() *QPaintEvent
}

func (ptr *QPaintEvent) QPaintEvent_PTR() *QPaintEvent {
	return ptr
}

func (ptr *QPaintEvent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QEvent_PTR().Pointer()
	}
	return nil
}

func (ptr *QPaintEvent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QEvent_PTR().SetPointer(p)
	}
}

func PointerFromQPaintEvent(ptr QPaintEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintEvent_PTR().Pointer()
	}
	return nil
}

func NewQPaintEventFromPointer(ptr unsafe.Pointer) (n *QPaintEvent) {
	n = new(QPaintEvent)
	n.SetPointer(ptr)
	return
}
func (ptr *QPaintEvent) DestroyQPaintEvent() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		qt.DisconnectAllSignals(ptr.Pointer(), "")
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQPaintEvent2(paintRect core.QRect_ITF) *QPaintEvent {
	tmpValue := NewQPaintEventFromPointer(C.QPaintEvent_NewQPaintEvent2(core.PointerFromQRect(paintRect)))
	qt.SetFinalizer(tmpValue, (*QPaintEvent).DestroyQPaintEvent)
	return tmpValue
}

func NewQPaintEvent(paintRegion QRegion_ITF) *QPaintEvent {
	tmpValue := NewQPaintEventFromPointer(C.QPaintEvent_NewQPaintEvent(PointerFromQRegion(paintRegion)))
	qt.SetFinalizer(tmpValue, (*QPaintEvent).DestroyQPaintEvent)
	return tmpValue
}

type QPainter struct {
	ptr unsafe.Pointer
}

type QPainter_ITF interface {
	QPainter_PTR() *QPainter
}

func (ptr *QPainter) QPainter_PTR() *QPainter {
	return ptr
}

func (ptr *QPainter) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPainter) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPainter(ptr QPainter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPainter_PTR().Pointer()
	}
	return nil
}

func NewQPainterFromPointer(ptr unsafe.Pointer) (n *QPainter) {
	n = new(QPainter)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QPainter__CompositionMode
//QPainter::CompositionMode
type QPainter__CompositionMode int64

const (
	QPainter__CompositionMode_SourceOver          QPainter__CompositionMode = QPainter__CompositionMode(0)
	QPainter__CompositionMode_DestinationOver     QPainter__CompositionMode = QPainter__CompositionMode(1)
	QPainter__CompositionMode_Clear               QPainter__CompositionMode = QPainter__CompositionMode(2)
	QPainter__CompositionMode_Source              QPainter__CompositionMode = QPainter__CompositionMode(3)
	QPainter__CompositionMode_Destination         QPainter__CompositionMode = QPainter__CompositionMode(4)
	QPainter__CompositionMode_SourceIn            QPainter__CompositionMode = QPainter__CompositionMode(5)
	QPainter__CompositionMode_DestinationIn       QPainter__CompositionMode = QPainter__CompositionMode(6)
	QPainter__CompositionMode_SourceOut           QPainter__CompositionMode = QPainter__CompositionMode(7)
	QPainter__CompositionMode_DestinationOut      QPainter__CompositionMode = QPainter__CompositionMode(8)
	QPainter__CompositionMode_SourceAtop          QPainter__CompositionMode = QPainter__CompositionMode(9)
	QPainter__CompositionMode_DestinationAtop     QPainter__CompositionMode = QPainter__CompositionMode(10)
	QPainter__CompositionMode_Xor                 QPainter__CompositionMode = QPainter__CompositionMode(11)
	QPainter__CompositionMode_Plus                QPainter__CompositionMode = QPainter__CompositionMode(12)
	QPainter__CompositionMode_Multiply            QPainter__CompositionMode = QPainter__CompositionMode(13)
	QPainter__CompositionMode_Screen              QPainter__CompositionMode = QPainter__CompositionMode(14)
	QPainter__CompositionMode_Overlay             QPainter__CompositionMode = QPainter__CompositionMode(15)
	QPainter__CompositionMode_Darken              QPainter__CompositionMode = QPainter__CompositionMode(16)
	QPainter__CompositionMode_Lighten             QPainter__CompositionMode = QPainter__CompositionMode(17)
	QPainter__CompositionMode_ColorDodge          QPainter__CompositionMode = QPainter__CompositionMode(18)
	QPainter__CompositionMode_ColorBurn           QPainter__CompositionMode = QPainter__CompositionMode(19)
	QPainter__CompositionMode_HardLight           QPainter__CompositionMode = QPainter__CompositionMode(20)
	QPainter__CompositionMode_SoftLight           QPainter__CompositionMode = QPainter__CompositionMode(21)
	QPainter__CompositionMode_Difference          QPainter__CompositionMode = QPainter__CompositionMode(22)
	QPainter__CompositionMode_Exclusion           QPainter__CompositionMode = QPainter__CompositionMode(23)
	QPainter__RasterOp_SourceOrDestination        QPainter__CompositionMode = QPainter__CompositionMode(24)
	QPainter__RasterOp_SourceAndDestination       QPainter__CompositionMode = QPainter__CompositionMode(25)
	QPainter__RasterOp_SourceXorDestination       QPainter__CompositionMode = QPainter__CompositionMode(26)
	QPainter__RasterOp_NotSourceAndNotDestination QPainter__CompositionMode = QPainter__CompositionMode(27)
	QPainter__RasterOp_NotSourceOrNotDestination  QPainter__CompositionMode = QPainter__CompositionMode(28)
	QPainter__RasterOp_NotSourceXorDestination    QPainter__CompositionMode = QPainter__CompositionMode(29)
	QPainter__RasterOp_NotSource                  QPainter__CompositionMode = QPainter__CompositionMode(30)
	QPainter__RasterOp_NotSourceAndDestination    QPainter__CompositionMode = QPainter__CompositionMode(31)
	QPainter__RasterOp_SourceAndNotDestination    QPainter__CompositionMode = QPainter__CompositionMode(32)
	QPainter__RasterOp_NotSourceOrDestination     QPainter__CompositionMode = QPainter__CompositionMode(33)
	QPainter__RasterOp_SourceOrNotDestination     QPainter__CompositionMode = QPainter__CompositionMode(34)
	QPainter__RasterOp_ClearDestination           QPainter__CompositionMode = QPainter__CompositionMode(35)
	QPainter__RasterOp_SetDestination             QPainter__CompositionMode = QPainter__CompositionMode(36)
	QPainter__RasterOp_NotDestination             QPainter__CompositionMode = QPainter__CompositionMode(37)
)

//go:generate stringer -type=QPainter__PixmapFragmentHint
//QPainter::PixmapFragmentHint
type QPainter__PixmapFragmentHint int64

const (
	QPainter__OpaqueHint QPainter__PixmapFragmentHint = QPainter__PixmapFragmentHint(0x01)
)

//go:generate stringer -type=QPainter__RenderHint
//QPainter::RenderHint
type QPainter__RenderHint int64

const (
	QPainter__Antialiasing            QPainter__RenderHint = QPainter__RenderHint(0x01)
	QPainter__TextAntialiasing        QPainter__RenderHint = QPainter__RenderHint(0x02)
	QPainter__SmoothPixmapTransform   QPainter__RenderHint = QPainter__RenderHint(0x04)
	QPainter__HighQualityAntialiasing QPainter__RenderHint = QPainter__RenderHint(0x08)
	QPainter__NonCosmeticDefaultPen   QPainter__RenderHint = QPainter__RenderHint(0x10)
	QPainter__Qt4CompatiblePainting   QPainter__RenderHint = QPainter__RenderHint(0x20)
)

func NewQPainter2(device QPaintDevice_ITF) *QPainter {
	tmpValue := NewQPainterFromPointer(C.QPainter_NewQPainter2(PointerFromQPaintDevice(device)))
	qt.SetFinalizer(tmpValue, (*QPainter).DestroyQPainter)
	return tmpValue
}

func NewQPainter() *QPainter {
	tmpValue := NewQPainterFromPointer(C.QPainter_NewQPainter())
	qt.SetFinalizer(tmpValue, (*QPainter).DestroyQPainter)
	return tmpValue
}

func (ptr *QPainter) SetWindow(rectangle core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QPainter_SetWindow(ptr.Pointer(), core.PointerFromQRect(rectangle))
	}
}

func (ptr *QPainter) SetWindow2(x int, y int, width int, height int) {
	if ptr.Pointer() != nil {
		C.QPainter_SetWindow2(ptr.Pointer(), C.int(int32(x)), C.int(int32(y)), C.int(int32(width)), C.int(int32(height)))
	}
}

func (ptr *QPainter) Translate2(offset core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QPainter_Translate2(ptr.Pointer(), core.PointerFromQPoint(offset))
	}
}

func (ptr *QPainter) Translate(offset core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QPainter_Translate(ptr.Pointer(), core.PointerFromQPointF(offset))
	}
}

func (ptr *QPainter) Translate3(dx float64, dy float64) {
	if ptr.Pointer() != nil {
		C.QPainter_Translate3(ptr.Pointer(), C.double(dx), C.double(dy))
	}
}

func (ptr *QPainter) DestroyQPainter() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPainter_DestroyQPainter(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPainter) Window() *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QPainter_Window(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QPainter) __drawLines_lines_atList6(i int) *core.QLine {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQLineFromPointer(C.QPainter___drawLines_lines_atList6(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QLine).DestroyQLine)
		return tmpValue
	}
	return nil
}

func (ptr *QPainter) __drawLines_lines_setList6(i core.QLine_ITF) {
	if ptr.Pointer() != nil {
		C.QPainter___drawLines_lines_setList6(ptr.Pointer(), core.PointerFromQLine(i))
	}
}

func (ptr *QPainter) __drawLines_lines_newList6() unsafe.Pointer {
	return C.QPainter___drawLines_lines_newList6(ptr.Pointer())
}

func (ptr *QPainter) __drawLines_lines_atList2(i int) *core.QLineF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQLineFFromPointer(C.QPainter___drawLines_lines_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QLineF).DestroyQLineF)
		return tmpValue
	}
	return nil
}

func (ptr *QPainter) __drawLines_lines_setList2(i core.QLineF_ITF) {
	if ptr.Pointer() != nil {
		C.QPainter___drawLines_lines_setList2(ptr.Pointer(), core.PointerFromQLineF(i))
	}
}

func (ptr *QPainter) __drawLines_lines_newList2() unsafe.Pointer {
	return C.QPainter___drawLines_lines_newList2(ptr.Pointer())
}

func (ptr *QPainter) __drawLines_pointPairs_atList8(i int) *core.QPoint {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFromPointer(C.QPainter___drawLines_pointPairs_atList8(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QPainter) __drawLines_pointPairs_setList8(i core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QPainter___drawLines_pointPairs_setList8(ptr.Pointer(), core.PointerFromQPoint(i))
	}
}

func (ptr *QPainter) __drawLines_pointPairs_newList8() unsafe.Pointer {
	return C.QPainter___drawLines_pointPairs_newList8(ptr.Pointer())
}

func (ptr *QPainter) __drawLines_pointPairs_atList4(i int) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QPainter___drawLines_pointPairs_atList4(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QPainter) __drawLines_pointPairs_setList4(i core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QPainter___drawLines_pointPairs_setList4(ptr.Pointer(), core.PointerFromQPointF(i))
	}
}

func (ptr *QPainter) __drawLines_pointPairs_newList4() unsafe.Pointer {
	return C.QPainter___drawLines_pointPairs_newList4(ptr.Pointer())
}

func (ptr *QPainter) __drawRects_rectangles_atList4(i int) *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QPainter___drawRects_rectangles_atList4(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QPainter) __drawRects_rectangles_setList4(i core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QPainter___drawRects_rectangles_setList4(ptr.Pointer(), core.PointerFromQRect(i))
	}
}

func (ptr *QPainter) __drawRects_rectangles_newList4() unsafe.Pointer {
	return C.QPainter___drawRects_rectangles_newList4(ptr.Pointer())
}

func (ptr *QPainter) __drawRects_rectangles_atList2(i int) *core.QRectF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFFromPointer(C.QPainter___drawRects_rectangles_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QPainter) __drawRects_rectangles_setList2(i core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QPainter___drawRects_rectangles_setList2(ptr.Pointer(), core.PointerFromQRectF(i))
	}
}

func (ptr *QPainter) __drawRects_rectangles_newList2() unsafe.Pointer {
	return C.QPainter___drawRects_rectangles_newList2(ptr.Pointer())
}

type QPainterPath struct {
	ptr unsafe.Pointer
}

type QPainterPath_ITF interface {
	QPainterPath_PTR() *QPainterPath
}

func (ptr *QPainterPath) QPainterPath_PTR() *QPainterPath {
	return ptr
}

func (ptr *QPainterPath) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QPainterPath) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQPainterPath(ptr QPainterPath_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPainterPath_PTR().Pointer()
	}
	return nil
}

func NewQPainterPathFromPointer(ptr unsafe.Pointer) (n *QPainterPath) {
	n = new(QPainterPath)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QPainterPath__ElementType
//QPainterPath::ElementType
type QPainterPath__ElementType int64

const (
	QPainterPath__MoveToElement      QPainterPath__ElementType = QPainterPath__ElementType(0)
	QPainterPath__LineToElement      QPainterPath__ElementType = QPainterPath__ElementType(1)
	QPainterPath__CurveToElement     QPainterPath__ElementType = QPainterPath__ElementType(2)
	QPainterPath__CurveToDataElement QPainterPath__ElementType = QPainterPath__ElementType(3)
)

func NewQPainterPath3(path QPainterPath_ITF) *QPainterPath {
	tmpValue := NewQPainterPathFromPointer(C.QPainterPath_NewQPainterPath3(PointerFromQPainterPath(path)))
	qt.SetFinalizer(tmpValue, (*QPainterPath).DestroyQPainterPath)
	return tmpValue
}

func NewQPainterPath() *QPainterPath {
	tmpValue := NewQPainterPathFromPointer(C.QPainterPath_NewQPainterPath())
	qt.SetFinalizer(tmpValue, (*QPainterPath).DestroyQPainterPath)
	return tmpValue
}

func NewQPainterPath2(startPoint core.QPointF_ITF) *QPainterPath {
	tmpValue := NewQPainterPathFromPointer(C.QPainterPath_NewQPainterPath2(core.PointerFromQPointF(startPoint)))
	qt.SetFinalizer(tmpValue, (*QPainterPath).DestroyQPainterPath)
	return tmpValue
}

func (ptr *QPainterPath) Translate2(offset core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QPainterPath_Translate2(ptr.Pointer(), core.PointerFromQPointF(offset))
	}
}

func (ptr *QPainterPath) Translate(dx float64, dy float64) {
	if ptr.Pointer() != nil {
		C.QPainterPath_Translate(ptr.Pointer(), C.double(dx), C.double(dy))
	}
}

func (ptr *QPainterPath) DestroyQPainterPath() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPainterPath_DestroyQPainterPath(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPainterPath) Contains3(p QPainterPath_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPainterPath_Contains3(ptr.Pointer(), PointerFromQPainterPath(p))) != 0
	}
	return false
}

func (ptr *QPainterPath) Contains(point core.QPointF_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPainterPath_Contains(ptr.Pointer(), core.PointerFromQPointF(point))) != 0
	}
	return false
}

func (ptr *QPainterPath) Contains2(rectangle core.QRectF_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QPainterPath_Contains2(ptr.Pointer(), core.PointerFromQRectF(rectangle))) != 0
	}
	return false
}

func (ptr *QPainterPath) Length() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QPainterPath_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPainterPath) __toFillPolygons_atList2(i int) *QPolygonF {
	if ptr.Pointer() != nil {
		tmpValue := NewQPolygonFFromPointer(C.QPainterPath___toFillPolygons_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
		return tmpValue
	}
	return nil
}

func (ptr *QPainterPath) __toFillPolygons_setList2(i QPolygonF_ITF) {
	if ptr.Pointer() != nil {
		C.QPainterPath___toFillPolygons_setList2(ptr.Pointer(), PointerFromQPolygonF(i))
	}
}

func (ptr *QPainterPath) __toFillPolygons_newList2() unsafe.Pointer {
	return C.QPainterPath___toFillPolygons_newList2(ptr.Pointer())
}

func (ptr *QPainterPath) __toFillPolygons_atList(i int) *QPolygonF {
	if ptr.Pointer() != nil {
		tmpValue := NewQPolygonFFromPointer(C.QPainterPath___toFillPolygons_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
		return tmpValue
	}
	return nil
}

func (ptr *QPainterPath) __toFillPolygons_setList(i QPolygonF_ITF) {
	if ptr.Pointer() != nil {
		C.QPainterPath___toFillPolygons_setList(ptr.Pointer(), PointerFromQPolygonF(i))
	}
}

func (ptr *QPainterPath) __toFillPolygons_newList() unsafe.Pointer {
	return C.QPainterPath___toFillPolygons_newList(ptr.Pointer())
}

func (ptr *QPainterPath) __toSubpathPolygons_atList2(i int) *QPolygonF {
	if ptr.Pointer() != nil {
		tmpValue := NewQPolygonFFromPointer(C.QPainterPath___toSubpathPolygons_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
		return tmpValue
	}
	return nil
}

func (ptr *QPainterPath) __toSubpathPolygons_setList2(i QPolygonF_ITF) {
	if ptr.Pointer() != nil {
		C.QPainterPath___toSubpathPolygons_setList2(ptr.Pointer(), PointerFromQPolygonF(i))
	}
}

func (ptr *QPainterPath) __toSubpathPolygons_newList2() unsafe.Pointer {
	return C.QPainterPath___toSubpathPolygons_newList2(ptr.Pointer())
}

func (ptr *QPainterPath) __toSubpathPolygons_atList(i int) *QPolygonF {
	if ptr.Pointer() != nil {
		tmpValue := NewQPolygonFFromPointer(C.QPainterPath___toSubpathPolygons_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
		return tmpValue
	}
	return nil
}

func (ptr *QPainterPath) __toSubpathPolygons_setList(i QPolygonF_ITF) {
	if ptr.Pointer() != nil {
		C.QPainterPath___toSubpathPolygons_setList(ptr.Pointer(), PointerFromQPolygonF(i))
	}
}

func (ptr *QPainterPath) __toSubpathPolygons_newList() unsafe.Pointer {
	return C.QPainterPath___toSubpathPolygons_newList(ptr.Pointer())
}

//go:generate stringer -type=QPalette__ColorGroup
//QPalette::ColorGroup
type QPalette__ColorGroup int64

const (
	QPalette__Active       QPalette__ColorGroup = QPalette__ColorGroup(0)
	QPalette__Disabled     QPalette__ColorGroup = QPalette__ColorGroup(1)
	QPalette__Inactive     QPalette__ColorGroup = QPalette__ColorGroup(2)
	QPalette__NColorGroups QPalette__ColorGroup = QPalette__ColorGroup(3)
	QPalette__Current      QPalette__ColorGroup = QPalette__ColorGroup(4)
	QPalette__All          QPalette__ColorGroup = QPalette__ColorGroup(5)
	QPalette__Normal       QPalette__ColorGroup = QPalette__ColorGroup(QPalette__Active)
)

//go:generate stringer -type=QPalette__ColorRole
//QPalette::ColorRole
type QPalette__ColorRole int64

var (
	QPalette__WindowText      QPalette__ColorRole = QPalette__ColorRole(0)
	QPalette__Button          QPalette__ColorRole = QPalette__ColorRole(1)
	QPalette__Light           QPalette__ColorRole = QPalette__ColorRole(2)
	QPalette__Midlight        QPalette__ColorRole = QPalette__ColorRole(3)
	QPalette__Dark            QPalette__ColorRole = QPalette__ColorRole(4)
	QPalette__Mid             QPalette__ColorRole = QPalette__ColorRole(5)
	QPalette__Text            QPalette__ColorRole = QPalette__ColorRole(6)
	QPalette__BrightText      QPalette__ColorRole = QPalette__ColorRole(7)
	QPalette__ButtonText      QPalette__ColorRole = QPalette__ColorRole(8)
	QPalette__Base            QPalette__ColorRole = QPalette__ColorRole(9)
	QPalette__Window          QPalette__ColorRole = QPalette__ColorRole(10)
	QPalette__Shadow          QPalette__ColorRole = QPalette__ColorRole(11)
	QPalette__Highlight       QPalette__ColorRole = QPalette__ColorRole(12)
	QPalette__HighlightedText QPalette__ColorRole = QPalette__ColorRole(13)
	QPalette__Link            QPalette__ColorRole = QPalette__ColorRole(14)
	QPalette__LinkVisited     QPalette__ColorRole = QPalette__ColorRole(15)
	QPalette__AlternateBase   QPalette__ColorRole = QPalette__ColorRole(16)
	QPalette__NoRole          QPalette__ColorRole = QPalette__ColorRole(17)
	QPalette__ToolTipBase     QPalette__ColorRole = QPalette__ColorRole(18)
	QPalette__ToolTipText     QPalette__ColorRole = QPalette__ColorRole(19)
	QPalette__NColorRoles     QPalette__ColorRole = QPalette__ColorRole(C.QPalette_NColorRoles_Type())
	QPalette__Foreground      QPalette__ColorRole = QPalette__ColorRole(QPalette__WindowText)
	QPalette__Background      QPalette__ColorRole = QPalette__ColorRole(QPalette__Window)
)

//go:generate stringer -type=QPdf__PathFlags
//QPdf::PathFlags
type QPdf__PathFlags int64

const (
	QPdf__ClipPath          QPdf__PathFlags = QPdf__PathFlags(0)
	QPdf__FillPath          QPdf__PathFlags = QPdf__PathFlags(1)
	QPdf__StrokePath        QPdf__PathFlags = QPdf__PathFlags(2)
	QPdf__FillAndStrokePath QPdf__PathFlags = QPdf__PathFlags(3)
)

//go:generate stringer -type=QPixelFormat__AlphaPosition
//QPixelFormat::AlphaPosition
type QPixelFormat__AlphaPosition int64

const (
	QPixelFormat__AtBeginning QPixelFormat__AlphaPosition = QPixelFormat__AlphaPosition(0)
	QPixelFormat__AtEnd       QPixelFormat__AlphaPosition = QPixelFormat__AlphaPosition(1)
)

//go:generate stringer -type=QPixelFormat__AlphaPremultiplied
//QPixelFormat::AlphaPremultiplied
type QPixelFormat__AlphaPremultiplied int64

const (
	QPixelFormat__NotPremultiplied QPixelFormat__AlphaPremultiplied = QPixelFormat__AlphaPremultiplied(0)
	QPixelFormat__Premultiplied    QPixelFormat__AlphaPremultiplied = QPixelFormat__AlphaPremultiplied(1)
)

//go:generate stringer -type=QPixelFormat__AlphaUsage
//QPixelFormat::AlphaUsage
type QPixelFormat__AlphaUsage int64

const (
	QPixelFormat__UsesAlpha    QPixelFormat__AlphaUsage = QPixelFormat__AlphaUsage(0)
	QPixelFormat__IgnoresAlpha QPixelFormat__AlphaUsage = QPixelFormat__AlphaUsage(1)
)

//go:generate stringer -type=QPixelFormat__ByteOrder
//QPixelFormat::ByteOrder
type QPixelFormat__ByteOrder int64

const (
	QPixelFormat__LittleEndian        QPixelFormat__ByteOrder = QPixelFormat__ByteOrder(0)
	QPixelFormat__BigEndian           QPixelFormat__ByteOrder = QPixelFormat__ByteOrder(1)
	QPixelFormat__CurrentSystemEndian QPixelFormat__ByteOrder = QPixelFormat__ByteOrder(2)
)

//go:generate stringer -type=QPixelFormat__ColorModel
//QPixelFormat::ColorModel
type QPixelFormat__ColorModel int64

const (
	QPixelFormat__RGB       QPixelFormat__ColorModel = QPixelFormat__ColorModel(0)
	QPixelFormat__BGR       QPixelFormat__ColorModel = QPixelFormat__ColorModel(1)
	QPixelFormat__Indexed   QPixelFormat__ColorModel = QPixelFormat__ColorModel(2)
	QPixelFormat__Grayscale QPixelFormat__ColorModel = QPixelFormat__ColorModel(3)
	QPixelFormat__CMYK      QPixelFormat__ColorModel = QPixelFormat__ColorModel(4)
	QPixelFormat__HSL       QPixelFormat__ColorModel = QPixelFormat__ColorModel(5)
	QPixelFormat__HSV       QPixelFormat__ColorModel = QPixelFormat__ColorModel(6)
	QPixelFormat__YUV       QPixelFormat__ColorModel = QPixelFormat__ColorModel(7)
	QPixelFormat__Alpha     QPixelFormat__ColorModel = QPixelFormat__ColorModel(8)
)

//go:generate stringer -type=QPixelFormat__TypeInterpretation
//QPixelFormat::TypeInterpretation
type QPixelFormat__TypeInterpretation int64

const (
	QPixelFormat__UnsignedInteger QPixelFormat__TypeInterpretation = QPixelFormat__TypeInterpretation(0)
	QPixelFormat__UnsignedShort   QPixelFormat__TypeInterpretation = QPixelFormat__TypeInterpretation(1)
	QPixelFormat__UnsignedByte    QPixelFormat__TypeInterpretation = QPixelFormat__TypeInterpretation(2)
	QPixelFormat__FloatingPoint   QPixelFormat__TypeInterpretation = QPixelFormat__TypeInterpretation(3)
)

//go:generate stringer -type=QPixelFormat__YUVLayout
//QPixelFormat::YUVLayout
type QPixelFormat__YUVLayout int64

const (
	QPixelFormat__YUV444   QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(0)
	QPixelFormat__YUV422   QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(1)
	QPixelFormat__YUV411   QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(2)
	QPixelFormat__YUV420P  QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(3)
	QPixelFormat__YUV420SP QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(4)
	QPixelFormat__YV12     QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(5)
	QPixelFormat__UYVY     QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(6)
	QPixelFormat__YUYV     QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(7)
	QPixelFormat__NV12     QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(8)
	QPixelFormat__NV21     QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(9)
	QPixelFormat__IMC1     QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(10)
	QPixelFormat__IMC2     QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(11)
	QPixelFormat__IMC3     QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(12)
	QPixelFormat__IMC4     QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(13)
	QPixelFormat__Y8       QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(14)
	QPixelFormat__Y16      QPixelFormat__YUVLayout = QPixelFormat__YUVLayout(15)
)

type QPixmap struct {
	QPaintDevice
}

type QPixmap_ITF interface {
	QPaintDevice_ITF
	QPixmap_PTR() *QPixmap
}

func (ptr *QPixmap) QPixmap_PTR() *QPixmap {
	return ptr
}

func (ptr *QPixmap) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintDevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QPixmap) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QPaintDevice_PTR().SetPointer(p)
	}
}

func PointerFromQPixmap(ptr QPixmap_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPixmap_PTR().Pointer()
	}
	return nil
}

func NewQPixmapFromPointer(ptr unsafe.Pointer) (n *QPixmap) {
	n = new(QPixmap)
	n.SetPointer(ptr)
	return
}
func NewQPixmap() *QPixmap {
	return NewQPixmapFromPointer(C.QPixmap_NewQPixmap())
}

func NewQPixmap7(pixmap QPixmap_ITF) *QPixmap {
	return NewQPixmapFromPointer(C.QPixmap_NewQPixmap7(PointerFromQPixmap(pixmap)))
}

func NewQPixmap4(size core.QSize_ITF) *QPixmap {
	return NewQPixmapFromPointer(C.QPixmap_NewQPixmap4(core.PointerFromQSize(size)))
}

func NewQPixmap5(fileName string, format string, flags core.Qt__ImageConversionFlag) *QPixmap {
	var fileNameC *C.char
	if fileName != "" {
		fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
	}
	var formatC *C.char
	if format != "" {
		formatC = C.CString(format)
		defer C.free(unsafe.Pointer(formatC))
	}
	return NewQPixmapFromPointer(C.QPixmap_NewQPixmap5(C.struct_QtGui_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, formatC, C.longlong(flags)))
}

func NewQPixmap3(width int, height int) *QPixmap {
	return NewQPixmapFromPointer(C.QPixmap_NewQPixmap3(C.int(int32(width)), C.int(int32(height))))
}

func (ptr *QPixmap) Load(fileName string, format string, flags core.Qt__ImageConversionFlag) bool {
	if ptr.Pointer() != nil {
		var fileNameC *C.char
		if fileName != "" {
			fileNameC = C.CString(fileName)
			defer C.free(unsafe.Pointer(fileNameC))
		}
		var formatC *C.char
		if format != "" {
			formatC = C.CString(format)
			defer C.free(unsafe.Pointer(formatC))
		}
		return int8(C.QPixmap_Load(ptr.Pointer(), C.struct_QtGui_PackedString{data: fileNameC, len: C.longlong(len(fileName))}, formatC, C.longlong(flags))) != 0
	}
	return false
}

func (ptr *QPixmap) DestroyQPixmap() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPixmap_DestroyQPixmap(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPixmap) Copy(rectangle core.QRect_ITF) *QPixmap {
	if ptr.Pointer() != nil {
		tmpValue := NewQPixmapFromPointer(C.QPixmap_Copy(ptr.Pointer(), core.PointerFromQRect(rectangle)))
		qt.SetFinalizer(tmpValue, (*QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

func (ptr *QPixmap) Copy2(x int, y int, width int, height int) *QPixmap {
	if ptr.Pointer() != nil {
		tmpValue := NewQPixmapFromPointer(C.QPixmap_Copy2(ptr.Pointer(), C.int(int32(x)), C.int(int32(y)), C.int(int32(width)), C.int(int32(height))))
		qt.SetFinalizer(tmpValue, (*QPixmap).DestroyQPixmap)
		return tmpValue
	}
	return nil
}

func (ptr *QPixmap) Size() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QPixmap_Size(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QPixmap) ToVariant() *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QPixmap_ToVariant(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQPixmap_PaintEngine
func callbackQPixmap_PaintEngine(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "paintEngine"); signal != nil {
		return PointerFromQPaintEngine((*(*func() *QPaintEngine)(signal))())
	}

	return PointerFromQPaintEngine(NewQPixmapFromPointer(ptr).PaintEngineDefault())
}

func (ptr *QPixmap) PaintEngine() *QPaintEngine {
	if ptr.Pointer() != nil {
		return NewQPaintEngineFromPointer(C.QPixmap_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPixmap) PaintEngineDefault() *QPaintEngine {
	if ptr.Pointer() != nil {
		return NewQPaintEngineFromPointer(C.QPixmap_PaintEngineDefault(ptr.Pointer()))
	}
	return nil
}

//go:generate stringer -type=QPlatformSurfaceEvent__SurfaceEventType
//QPlatformSurfaceEvent::SurfaceEventType
type QPlatformSurfaceEvent__SurfaceEventType int64

const (
	QPlatformSurfaceEvent__SurfaceCreated            QPlatformSurfaceEvent__SurfaceEventType = QPlatformSurfaceEvent__SurfaceEventType(0)
	QPlatformSurfaceEvent__SurfaceAboutToBeDestroyed QPlatformSurfaceEvent__SurfaceEventType = QPlatformSurfaceEvent__SurfaceEventType(1)
)

type QPolygon struct {
	core.QVector
}

type QPolygon_ITF interface {
	core.QVector_ITF
	QPolygon_PTR() *QPolygon
}

func (ptr *QPolygon) QPolygon_PTR() *QPolygon {
	return ptr
}

func (ptr *QPolygon) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QVector_PTR().Pointer()
	}
	return nil
}

func (ptr *QPolygon) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QVector_PTR().SetPointer(p)
	}
}

func PointerFromQPolygon(ptr QPolygon_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPolygon_PTR().Pointer()
	}
	return nil
}

func NewQPolygonFromPointer(ptr unsafe.Pointer) (n *QPolygon) {
	n = new(QPolygon)
	n.SetPointer(ptr)
	return
}
func NewQPolygon() *QPolygon {
	tmpValue := NewQPolygonFromPointer(C.QPolygon_NewQPolygon())
	qt.SetFinalizer(tmpValue, (*QPolygon).DestroyQPolygon)
	return tmpValue
}

func NewQPolygon8(other QPolygon_ITF) *QPolygon {
	tmpValue := NewQPolygonFromPointer(C.QPolygon_NewQPolygon8(PointerFromQPolygon(other)))
	qt.SetFinalizer(tmpValue, (*QPolygon).DestroyQPolygon)
	return tmpValue
}

func NewQPolygon4(v []*core.QPoint) *QPolygon {
	tmpValue := NewQPolygonFromPointer(C.QPolygon_NewQPolygon4(func() unsafe.Pointer {
		tmpList := NewQPolygonFromPointer(NewQPolygonFromPointer(nil).__QPolygon_v_newList4())
		for _, v := range v {
			tmpList.__QPolygon_v_setList4(v)
		}
		return tmpList.Pointer()
	}()))
	qt.SetFinalizer(tmpValue, (*QPolygon).DestroyQPolygon)
	return tmpValue
}

func NewQPolygon7(polygon QPolygon_ITF) *QPolygon {
	tmpValue := NewQPolygonFromPointer(C.QPolygon_NewQPolygon7(PointerFromQPolygon(polygon)))
	qt.SetFinalizer(tmpValue, (*QPolygon).DestroyQPolygon)
	return tmpValue
}

func NewQPolygon5(rectangle core.QRect_ITF, closed bool) *QPolygon {
	tmpValue := NewQPolygonFromPointer(C.QPolygon_NewQPolygon5(core.PointerFromQRect(rectangle), C.char(int8(qt.GoBoolToInt(closed)))))
	qt.SetFinalizer(tmpValue, (*QPolygon).DestroyQPolygon)
	return tmpValue
}

func NewQPolygon3(points []*core.QPoint) *QPolygon {
	tmpValue := NewQPolygonFromPointer(C.QPolygon_NewQPolygon3(func() unsafe.Pointer {
		tmpList := NewQPolygonFromPointer(NewQPolygonFromPointer(nil).__QPolygon_points_newList3())
		for _, v := range points {
			tmpList.__QPolygon_points_setList3(v)
		}
		return tmpList.Pointer()
	}()))
	qt.SetFinalizer(tmpValue, (*QPolygon).DestroyQPolygon)
	return tmpValue
}

func NewQPolygon2(size int) *QPolygon {
	tmpValue := NewQPolygonFromPointer(C.QPolygon_NewQPolygon2(C.int(int32(size))))
	qt.SetFinalizer(tmpValue, (*QPolygon).DestroyQPolygon)
	return tmpValue
}

func (ptr *QPolygon) SetPoint2(index int, point core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon_SetPoint2(ptr.Pointer(), C.int(int32(index)), core.PointerFromQPoint(point))
	}
}

func (ptr *QPolygon) SetPoint(index int, x int, y int) {
	if ptr.Pointer() != nil {
		C.QPolygon_SetPoint(ptr.Pointer(), C.int(int32(index)), C.int(int32(x)), C.int(int32(y)))
	}
}

func (ptr *QPolygon) Translate2(offset core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon_Translate2(ptr.Pointer(), core.PointerFromQPoint(offset))
	}
}

func (ptr *QPolygon) Translate(dx int, dy int) {
	if ptr.Pointer() != nil {
		C.QPolygon_Translate(ptr.Pointer(), C.int(int32(dx)), C.int(int32(dy)))
	}
}

func (ptr *QPolygon) DestroyQPolygon() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPolygon_DestroyQPolygon(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPolygon) Point2(index int) *core.QPoint {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFromPointer(C.QPolygon_Point2(ptr.Pointer(), C.int(int32(index))))
		qt.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) Point(index int, x int, y int) {
	if ptr.Pointer() != nil {
		C.QPolygon_Point(ptr.Pointer(), C.int(int32(index)), C.int(int32(x)), C.int(int32(y)))
	}
}

func (ptr *QPolygon) __QPolygon_v_atList4(i int) *core.QPoint {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFromPointer(C.QPolygon___QPolygon_v_atList4(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __QPolygon_v_setList4(i core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___QPolygon_v_setList4(ptr.Pointer(), core.PointerFromQPoint(i))
	}
}

func (ptr *QPolygon) __QPolygon_v_newList4() unsafe.Pointer {
	return C.QPolygon___QPolygon_v_newList4(ptr.Pointer())
}

func (ptr *QPolygon) __QPolygon_points_atList3(i int) *core.QPoint {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFromPointer(C.QPolygon___QPolygon_points_atList3(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __QPolygon_points_setList3(i core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___QPolygon_points_setList3(ptr.Pointer(), core.PointerFromQPoint(i))
	}
}

func (ptr *QPolygon) __QPolygon_points_newList3() unsafe.Pointer {
	return C.QPolygon___QPolygon_points_newList3(ptr.Pointer())
}

func (ptr *QPolygon) __QVector_other_atList5(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygon___QVector_other_atList5(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __QVector_other_setList5(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___QVector_other_setList5(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygon) __QVector_other_newList5() unsafe.Pointer {
	return C.QPolygon___QVector_other_newList5(ptr.Pointer())
}

func (ptr *QPolygon) __QVector_other_atList4(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygon___QVector_other_atList4(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __QVector_other_setList4(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___QVector_other_setList4(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygon) __QVector_other_newList4() unsafe.Pointer {
	return C.QPolygon___QVector_other_newList4(ptr.Pointer())
}

func (ptr *QPolygon) __fill_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygon___fill_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __fill_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___fill_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygon) __fill_newList() unsafe.Pointer {
	return C.QPolygon___fill_newList(ptr.Pointer())
}

func (ptr *QPolygon) __fromList_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygon___fromList_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __fromList_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___fromList_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygon) __fromList_newList() unsafe.Pointer {
	return C.QPolygon___fromList_newList(ptr.Pointer())
}

func (ptr *QPolygon) __fromList_list_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygon___fromList_list_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __fromList_list_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___fromList_list_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygon) __fromList_list_newList() unsafe.Pointer {
	return C.QPolygon___fromList_list_newList(ptr.Pointer())
}

func (ptr *QPolygon) __fromStdVector_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygon___fromStdVector_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __fromStdVector_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___fromStdVector_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygon) __fromStdVector_newList() unsafe.Pointer {
	return C.QPolygon___fromStdVector_newList(ptr.Pointer())
}

func (ptr *QPolygon) __append_value_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygon___append_value_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __append_value_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___append_value_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygon) __append_value_newList3() unsafe.Pointer {
	return C.QPolygon___append_value_newList3(ptr.Pointer())
}

func (ptr *QPolygon) __swap_other_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygon___swap_other_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __swap_other_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___swap_other_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygon) __swap_other_newList() unsafe.Pointer {
	return C.QPolygon___swap_other_newList(ptr.Pointer())
}

func (ptr *QPolygon) __toList_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygon___toList_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __toList_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___toList_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygon) __toList_newList() unsafe.Pointer {
	return C.QPolygon___toList_newList(ptr.Pointer())
}

func (ptr *QPolygon) __mid_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygon___mid_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygon) __mid_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygon___mid_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygon) __mid_newList() unsafe.Pointer {
	return C.QPolygon___mid_newList(ptr.Pointer())
}

type QPolygonF struct {
	core.QVector
}

type QPolygonF_ITF interface {
	core.QVector_ITF
	QPolygonF_PTR() *QPolygonF
}

func (ptr *QPolygonF) QPolygonF_PTR() *QPolygonF {
	return ptr
}

func (ptr *QPolygonF) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QVector_PTR().Pointer()
	}
	return nil
}

func (ptr *QPolygonF) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QVector_PTR().SetPointer(p)
	}
}

func PointerFromQPolygonF(ptr QPolygonF_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPolygonF_PTR().Pointer()
	}
	return nil
}

func NewQPolygonFFromPointer(ptr unsafe.Pointer) (n *QPolygonF) {
	n = new(QPolygonF)
	n.SetPointer(ptr)
	return
}
func NewQPolygonF() *QPolygonF {
	tmpValue := NewQPolygonFFromPointer(C.QPolygonF_NewQPolygonF())
	qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
	return tmpValue
}

func NewQPolygonF8(other QPolygonF_ITF) *QPolygonF {
	tmpValue := NewQPolygonFFromPointer(C.QPolygonF_NewQPolygonF8(PointerFromQPolygonF(other)))
	qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
	return tmpValue
}

func NewQPolygonF4(v []*core.QPointF) *QPolygonF {
	tmpValue := NewQPolygonFFromPointer(C.QPolygonF_NewQPolygonF4(func() unsafe.Pointer {
		tmpList := NewQPolygonFFromPointer(NewQPolygonFFromPointer(nil).__QPolygonF_v_newList4())
		for _, v := range v {
			tmpList.__QPolygonF_v_setList4(v)
		}
		return tmpList.Pointer()
	}()))
	qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
	return tmpValue
}

func NewQPolygonF6(polygon QPolygon_ITF) *QPolygonF {
	tmpValue := NewQPolygonFFromPointer(C.QPolygonF_NewQPolygonF6(PointerFromQPolygon(polygon)))
	qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
	return tmpValue
}

func NewQPolygonF7(polygon QPolygonF_ITF) *QPolygonF {
	tmpValue := NewQPolygonFFromPointer(C.QPolygonF_NewQPolygonF7(PointerFromQPolygonF(polygon)))
	qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
	return tmpValue
}

func NewQPolygonF5(rectangle core.QRectF_ITF) *QPolygonF {
	tmpValue := NewQPolygonFFromPointer(C.QPolygonF_NewQPolygonF5(core.PointerFromQRectF(rectangle)))
	qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
	return tmpValue
}

func NewQPolygonF3(points []*core.QPointF) *QPolygonF {
	tmpValue := NewQPolygonFFromPointer(C.QPolygonF_NewQPolygonF3(func() unsafe.Pointer {
		tmpList := NewQPolygonFFromPointer(NewQPolygonFFromPointer(nil).__QPolygonF_points_newList3())
		for _, v := range points {
			tmpList.__QPolygonF_points_setList3(v)
		}
		return tmpList.Pointer()
	}()))
	qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
	return tmpValue
}

func NewQPolygonF2(size int) *QPolygonF {
	tmpValue := NewQPolygonFFromPointer(C.QPolygonF_NewQPolygonF2(C.int(int32(size))))
	qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
	return tmpValue
}

func (ptr *QPolygonF) Translate(offset core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF_Translate(ptr.Pointer(), core.PointerFromQPointF(offset))
	}
}

func (ptr *QPolygonF) Translate2(dx float64, dy float64) {
	if ptr.Pointer() != nil {
		C.QPolygonF_Translate2(ptr.Pointer(), C.double(dx), C.double(dy))
	}
}

func (ptr *QPolygonF) DestroyQPolygonF() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QPolygonF_DestroyQPolygonF(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPolygonF) __QPolygonF_v_atList4(i int) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QPolygonF___QPolygonF_v_atList4(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __QPolygonF_v_setList4(i core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___QPolygonF_v_setList4(ptr.Pointer(), core.PointerFromQPointF(i))
	}
}

func (ptr *QPolygonF) __QPolygonF_v_newList4() unsafe.Pointer {
	return C.QPolygonF___QPolygonF_v_newList4(ptr.Pointer())
}

func (ptr *QPolygonF) __QPolygonF_points_atList3(i int) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QPolygonF___QPolygonF_points_atList3(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __QPolygonF_points_setList3(i core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___QPolygonF_points_setList3(ptr.Pointer(), core.PointerFromQPointF(i))
	}
}

func (ptr *QPolygonF) __QPolygonF_points_newList3() unsafe.Pointer {
	return C.QPolygonF___QPolygonF_points_newList3(ptr.Pointer())
}

func (ptr *QPolygonF) __QVector_other_atList5(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygonF___QVector_other_atList5(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __QVector_other_setList5(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___QVector_other_setList5(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygonF) __QVector_other_newList5() unsafe.Pointer {
	return C.QPolygonF___QVector_other_newList5(ptr.Pointer())
}

func (ptr *QPolygonF) __QVector_other_atList4(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygonF___QVector_other_atList4(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __QVector_other_setList4(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___QVector_other_setList4(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygonF) __QVector_other_newList4() unsafe.Pointer {
	return C.QPolygonF___QVector_other_newList4(ptr.Pointer())
}

func (ptr *QPolygonF) __fill_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygonF___fill_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __fill_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___fill_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygonF) __fill_newList() unsafe.Pointer {
	return C.QPolygonF___fill_newList(ptr.Pointer())
}

func (ptr *QPolygonF) __fromList_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygonF___fromList_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __fromList_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___fromList_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygonF) __fromList_newList() unsafe.Pointer {
	return C.QPolygonF___fromList_newList(ptr.Pointer())
}

func (ptr *QPolygonF) __fromList_list_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygonF___fromList_list_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __fromList_list_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___fromList_list_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygonF) __fromList_list_newList() unsafe.Pointer {
	return C.QPolygonF___fromList_list_newList(ptr.Pointer())
}

func (ptr *QPolygonF) __fromStdVector_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygonF___fromStdVector_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __fromStdVector_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___fromStdVector_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygonF) __fromStdVector_newList() unsafe.Pointer {
	return C.QPolygonF___fromStdVector_newList(ptr.Pointer())
}

func (ptr *QPolygonF) __append_value_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygonF___append_value_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __append_value_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___append_value_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygonF) __append_value_newList3() unsafe.Pointer {
	return C.QPolygonF___append_value_newList3(ptr.Pointer())
}

func (ptr *QPolygonF) __swap_other_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygonF___swap_other_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __swap_other_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___swap_other_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygonF) __swap_other_newList() unsafe.Pointer {
	return C.QPolygonF___swap_other_newList(ptr.Pointer())
}

func (ptr *QPolygonF) __toList_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygonF___toList_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __toList_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___toList_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygonF) __toList_newList() unsafe.Pointer {
	return C.QPolygonF___toList_newList(ptr.Pointer())
}

func (ptr *QPolygonF) __mid_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QPolygonF___mid_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QPolygonF) __mid_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QPolygonF___mid_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QPolygonF) __mid_newList() unsafe.Pointer {
	return C.QPolygonF___mid_newList(ptr.Pointer())
}

//go:generate stringer -type=QRawFont__AntialiasingType
//QRawFont::AntialiasingType
type QRawFont__AntialiasingType int64

const (
	QRawFont__PixelAntialiasing    QRawFont__AntialiasingType = QRawFont__AntialiasingType(0)
	QRawFont__SubPixelAntialiasing QRawFont__AntialiasingType = QRawFont__AntialiasingType(1)
)

//go:generate stringer -type=QRawFont__LayoutFlag
//QRawFont::LayoutFlag
type QRawFont__LayoutFlag int64

const (
	QRawFont__SeparateAdvances QRawFont__LayoutFlag = QRawFont__LayoutFlag(0)
	QRawFont__KernedAdvances   QRawFont__LayoutFlag = QRawFont__LayoutFlag(1)
	QRawFont__UseDesignMetrics QRawFont__LayoutFlag = QRawFont__LayoutFlag(2)
)

type QRegion struct {
	ptr unsafe.Pointer
}

type QRegion_ITF interface {
	QRegion_PTR() *QRegion
}

func (ptr *QRegion) QRegion_PTR() *QRegion {
	return ptr
}

func (ptr *QRegion) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QRegion) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQRegion(ptr QRegion_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRegion_PTR().Pointer()
	}
	return nil
}

func NewQRegionFromPointer(ptr unsafe.Pointer) (n *QRegion) {
	n = new(QRegion)
	n.SetPointer(ptr)
	return
}
func (ptr *QRegion) DestroyQRegion() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QRegion__RegionType
//QRegion::RegionType
type QRegion__RegionType int64

const (
	QRegion__Rectangle QRegion__RegionType = QRegion__RegionType(0)
	QRegion__Ellipse   QRegion__RegionType = QRegion__RegionType(1)
)

func NewQRegion() *QRegion {
	tmpValue := NewQRegionFromPointer(C.QRegion_NewQRegion())
	qt.SetFinalizer(tmpValue, (*QRegion).DestroyQRegion)
	return tmpValue
}

func NewQRegion6(other QRegion_ITF) *QRegion {
	tmpValue := NewQRegionFromPointer(C.QRegion_NewQRegion6(PointerFromQRegion(other)))
	qt.SetFinalizer(tmpValue, (*QRegion).DestroyQRegion)
	return tmpValue
}

func NewQRegion7(bm QBitmap_ITF) *QRegion {
	tmpValue := NewQRegionFromPointer(C.QRegion_NewQRegion7(PointerFromQBitmap(bm)))
	qt.SetFinalizer(tmpValue, (*QRegion).DestroyQRegion)
	return tmpValue
}

func NewQRegion4(a QPolygon_ITF, fillRule core.Qt__FillRule) *QRegion {
	tmpValue := NewQRegionFromPointer(C.QRegion_NewQRegion4(PointerFromQPolygon(a), C.longlong(fillRule)))
	qt.SetFinalizer(tmpValue, (*QRegion).DestroyQRegion)
	return tmpValue
}

func NewQRegion3(r core.QRect_ITF, t QRegion__RegionType) *QRegion {
	tmpValue := NewQRegionFromPointer(C.QRegion_NewQRegion3(core.PointerFromQRect(r), C.longlong(t)))
	qt.SetFinalizer(tmpValue, (*QRegion).DestroyQRegion)
	return tmpValue
}

func NewQRegion5(r QRegion_ITF) *QRegion {
	tmpValue := NewQRegionFromPointer(C.QRegion_NewQRegion5(PointerFromQRegion(r)))
	qt.SetFinalizer(tmpValue, (*QRegion).DestroyQRegion)
	return tmpValue
}

func NewQRegion2(x int, y int, w int, h int, t QRegion__RegionType) *QRegion {
	tmpValue := NewQRegionFromPointer(C.QRegion_NewQRegion2(C.int(int32(x)), C.int(int32(y)), C.int(int32(w)), C.int(int32(h)), C.longlong(t)))
	qt.SetFinalizer(tmpValue, (*QRegion).DestroyQRegion)
	return tmpValue
}

func (ptr *QRegion) Translate2(point core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QRegion_Translate2(ptr.Pointer(), core.PointerFromQPoint(point))
	}
}

func (ptr *QRegion) Translate(dx int, dy int) {
	if ptr.Pointer() != nil {
		C.QRegion_Translate(ptr.Pointer(), C.int(int32(dx)), C.int(int32(dy)))
	}
}

func (ptr *QRegion) Contains(p core.QPoint_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRegion_Contains(ptr.Pointer(), core.PointerFromQPoint(p))) != 0
	}
	return false
}

func (ptr *QRegion) Contains2(r core.QRect_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QRegion_Contains2(ptr.Pointer(), core.PointerFromQRect(r))) != 0
	}
	return false
}

func (ptr *QRegion) __rects_atList(i int) *core.QRect {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQRectFromPointer(C.QRegion___rects_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QRegion) __rects_setList(i core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QRegion___rects_setList(ptr.Pointer(), core.PointerFromQRect(i))
	}
}

func (ptr *QRegion) __rects_newList() unsafe.Pointer {
	return C.QRegion___rects_newList(ptr.Pointer())
}

type QScreen struct {
	core.QObject
}

type QScreen_ITF interface {
	core.QObject_ITF
	QScreen_PTR() *QScreen
}

func (ptr *QScreen) QScreen_PTR() *QScreen {
	return ptr
}

func (ptr *QScreen) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QScreen) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQScreen(ptr QScreen_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScreen_PTR().Pointer()
	}
	return nil
}

func NewQScreenFromPointer(ptr unsafe.Pointer) (n *QScreen) {
	n = new(QScreen)
	n.SetPointer(ptr)
	return
}
func (ptr *QScreen) DestroyQScreen() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QScreen_DestroyQScreen(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScreen) Size() *core.QSize {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQSizeFromPointer(C.QScreen_Size(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QScreen) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QScreen_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScreen) Depth() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QScreen_Depth(ptr.Pointer())))
	}
	return 0
}

func (ptr *QScreen) __virtualSiblings_atList(i int) *QScreen {
	if ptr.Pointer() != nil {
		tmpValue := NewQScreenFromPointer(C.QScreen___virtualSiblings_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScreen) __virtualSiblings_setList(i QScreen_ITF) {
	if ptr.Pointer() != nil {
		C.QScreen___virtualSiblings_setList(ptr.Pointer(), PointerFromQScreen(i))
	}
}

func (ptr *QScreen) __virtualSiblings_newList() unsafe.Pointer {
	return C.QScreen___virtualSiblings_newList(ptr.Pointer())
}

func (ptr *QScreen) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QScreen___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QScreen) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QScreen___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QScreen) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QScreen___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QScreen) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QScreen___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScreen) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScreen___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScreen) __findChildren_newList2() unsafe.Pointer {
	return C.QScreen___findChildren_newList2(ptr.Pointer())
}

func (ptr *QScreen) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QScreen___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScreen) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScreen___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScreen) __findChildren_newList3() unsafe.Pointer {
	return C.QScreen___findChildren_newList3(ptr.Pointer())
}

func (ptr *QScreen) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QScreen___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScreen) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScreen___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScreen) __findChildren_newList() unsafe.Pointer {
	return C.QScreen___findChildren_newList(ptr.Pointer())
}

func (ptr *QScreen) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QScreen___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QScreen) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QScreen___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QScreen) __children_newList() unsafe.Pointer {
	return C.QScreen___children_newList(ptr.Pointer())
}

//export callbackQScreen_ObjectNameChanged
func callbackQScreen_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtGui_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//go:generate stringer -type=QScrollEvent__ScrollState
//QScrollEvent::ScrollState
type QScrollEvent__ScrollState int64

const (
	QScrollEvent__ScrollStarted  QScrollEvent__ScrollState = QScrollEvent__ScrollState(0)
	QScrollEvent__ScrollUpdated  QScrollEvent__ScrollState = QScrollEvent__ScrollState(1)
	QScrollEvent__ScrollFinished QScrollEvent__ScrollState = QScrollEvent__ScrollState(2)
)

//go:generate stringer -type=QSessionManager__RestartHint
//QSessionManager::RestartHint
type QSessionManager__RestartHint int64

const (
	QSessionManager__RestartIfRunning   QSessionManager__RestartHint = QSessionManager__RestartHint(0)
	QSessionManager__RestartAnyway      QSessionManager__RestartHint = QSessionManager__RestartHint(1)
	QSessionManager__RestartImmediately QSessionManager__RestartHint = QSessionManager__RestartHint(2)
	QSessionManager__RestartNever       QSessionManager__RestartHint = QSessionManager__RestartHint(3)
)

//go:generate stringer -type=QStandardItem__ItemType
//QStandardItem::ItemType
type QStandardItem__ItemType int64

const (
	QStandardItem__Type     QStandardItem__ItemType = QStandardItem__ItemType(0)
	QStandardItem__UserType QStandardItem__ItemType = QStandardItem__ItemType(1000)
)

//go:generate stringer -type=QStaticText__PerformanceHint
//QStaticText::PerformanceHint
type QStaticText__PerformanceHint int64

const (
	QStaticText__ModerateCaching   QStaticText__PerformanceHint = QStaticText__PerformanceHint(0)
	QStaticText__AggressiveCaching QStaticText__PerformanceHint = QStaticText__PerformanceHint(1)
)

//go:generate stringer -type=QSurface__SurfaceClass
//QSurface::SurfaceClass
type QSurface__SurfaceClass int64

const (
	QSurface__Window    QSurface__SurfaceClass = QSurface__SurfaceClass(0)
	QSurface__Offscreen QSurface__SurfaceClass = QSurface__SurfaceClass(1)
)

//go:generate stringer -type=QSurface__SurfaceType
//QSurface::SurfaceType
type QSurface__SurfaceType int64

const (
	QSurface__RasterSurface   QSurface__SurfaceType = QSurface__SurfaceType(0)
	QSurface__OpenGLSurface   QSurface__SurfaceType = QSurface__SurfaceType(1)
	QSurface__RasterGLSurface QSurface__SurfaceType = QSurface__SurfaceType(2)
	QSurface__OpenVGSurface   QSurface__SurfaceType = QSurface__SurfaceType(3)
)

//go:generate stringer -type=QSurfaceFormat__FormatOption
//QSurfaceFormat::FormatOption
type QSurfaceFormat__FormatOption int64

const (
	QSurfaceFormat__StereoBuffers       QSurfaceFormat__FormatOption = QSurfaceFormat__FormatOption(0x0001)
	QSurfaceFormat__DebugContext        QSurfaceFormat__FormatOption = QSurfaceFormat__FormatOption(0x0002)
	QSurfaceFormat__DeprecatedFunctions QSurfaceFormat__FormatOption = QSurfaceFormat__FormatOption(0x0004)
	QSurfaceFormat__ResetNotification   QSurfaceFormat__FormatOption = QSurfaceFormat__FormatOption(0x0008)
)

//go:generate stringer -type=QSurfaceFormat__OpenGLContextProfile
//QSurfaceFormat::OpenGLContextProfile
type QSurfaceFormat__OpenGLContextProfile int64

const (
	QSurfaceFormat__NoProfile            QSurfaceFormat__OpenGLContextProfile = QSurfaceFormat__OpenGLContextProfile(0)
	QSurfaceFormat__CoreProfile          QSurfaceFormat__OpenGLContextProfile = QSurfaceFormat__OpenGLContextProfile(1)
	QSurfaceFormat__CompatibilityProfile QSurfaceFormat__OpenGLContextProfile = QSurfaceFormat__OpenGLContextProfile(2)
)

//go:generate stringer -type=QSurfaceFormat__RenderableType
//QSurfaceFormat::RenderableType
type QSurfaceFormat__RenderableType int64

const (
	QSurfaceFormat__DefaultRenderableType QSurfaceFormat__RenderableType = QSurfaceFormat__RenderableType(0x0)
	QSurfaceFormat__OpenGL                QSurfaceFormat__RenderableType = QSurfaceFormat__RenderableType(0x1)
	QSurfaceFormat__OpenGLES              QSurfaceFormat__RenderableType = QSurfaceFormat__RenderableType(0x2)
	QSurfaceFormat__OpenVG                QSurfaceFormat__RenderableType = QSurfaceFormat__RenderableType(0x4)
)

//go:generate stringer -type=QSurfaceFormat__SwapBehavior
//QSurfaceFormat::SwapBehavior
type QSurfaceFormat__SwapBehavior int64

const (
	QSurfaceFormat__DefaultSwapBehavior QSurfaceFormat__SwapBehavior = QSurfaceFormat__SwapBehavior(0)
	QSurfaceFormat__SingleBuffer        QSurfaceFormat__SwapBehavior = QSurfaceFormat__SwapBehavior(1)
	QSurfaceFormat__DoubleBuffer        QSurfaceFormat__SwapBehavior = QSurfaceFormat__SwapBehavior(2)
	QSurfaceFormat__TripleBuffer        QSurfaceFormat__SwapBehavior = QSurfaceFormat__SwapBehavior(3)
)

//go:generate stringer -type=QTabletEvent__PointerType
//QTabletEvent::PointerType
type QTabletEvent__PointerType int64

const (
	QTabletEvent__UnknownPointer QTabletEvent__PointerType = QTabletEvent__PointerType(0)
	QTabletEvent__Pen            QTabletEvent__PointerType = QTabletEvent__PointerType(1)
	QTabletEvent__Cursor         QTabletEvent__PointerType = QTabletEvent__PointerType(2)
	QTabletEvent__Eraser         QTabletEvent__PointerType = QTabletEvent__PointerType(3)
)

//go:generate stringer -type=QTabletEvent__TabletDevice
//QTabletEvent::TabletDevice
type QTabletEvent__TabletDevice int64

const (
	QTabletEvent__NoDevice       QTabletEvent__TabletDevice = QTabletEvent__TabletDevice(0)
	QTabletEvent__Puck           QTabletEvent__TabletDevice = QTabletEvent__TabletDevice(1)
	QTabletEvent__Stylus         QTabletEvent__TabletDevice = QTabletEvent__TabletDevice(2)
	QTabletEvent__Airbrush       QTabletEvent__TabletDevice = QTabletEvent__TabletDevice(3)
	QTabletEvent__FourDMouse     QTabletEvent__TabletDevice = QTabletEvent__TabletDevice(4)
	QTabletEvent__XFreeEraser    QTabletEvent__TabletDevice = QTabletEvent__TabletDevice(5)
	QTabletEvent__RotationStylus QTabletEvent__TabletDevice = QTabletEvent__TabletDevice(6)
)

//go:generate stringer -type=QTextBlockFormat__LineHeightTypes
//QTextBlockFormat::LineHeightTypes
type QTextBlockFormat__LineHeightTypes int64

const (
	QTextBlockFormat__SingleHeight       QTextBlockFormat__LineHeightTypes = QTextBlockFormat__LineHeightTypes(0)
	QTextBlockFormat__ProportionalHeight QTextBlockFormat__LineHeightTypes = QTextBlockFormat__LineHeightTypes(1)
	QTextBlockFormat__FixedHeight        QTextBlockFormat__LineHeightTypes = QTextBlockFormat__LineHeightTypes(2)
	QTextBlockFormat__MinimumHeight      QTextBlockFormat__LineHeightTypes = QTextBlockFormat__LineHeightTypes(3)
	QTextBlockFormat__LineDistanceHeight QTextBlockFormat__LineHeightTypes = QTextBlockFormat__LineHeightTypes(4)
)

//go:generate stringer -type=QTextCharFormat__FontPropertiesInheritanceBehavior
//QTextCharFormat::FontPropertiesInheritanceBehavior
type QTextCharFormat__FontPropertiesInheritanceBehavior int64

const (
	QTextCharFormat__FontPropertiesSpecifiedOnly QTextCharFormat__FontPropertiesInheritanceBehavior = QTextCharFormat__FontPropertiesInheritanceBehavior(0)
	QTextCharFormat__FontPropertiesAll           QTextCharFormat__FontPropertiesInheritanceBehavior = QTextCharFormat__FontPropertiesInheritanceBehavior(1)
)

//go:generate stringer -type=QTextCharFormat__UnderlineStyle
//QTextCharFormat::UnderlineStyle
type QTextCharFormat__UnderlineStyle int64

var (
	QTextCharFormat__NoUnderline         QTextCharFormat__UnderlineStyle = QTextCharFormat__UnderlineStyle(0)
	QTextCharFormat__SingleUnderline     QTextCharFormat__UnderlineStyle = QTextCharFormat__UnderlineStyle(1)
	QTextCharFormat__DashUnderline       QTextCharFormat__UnderlineStyle = QTextCharFormat__UnderlineStyle(2)
	QTextCharFormat__DotLine             QTextCharFormat__UnderlineStyle = QTextCharFormat__UnderlineStyle(3)
	QTextCharFormat__DashDotLine         QTextCharFormat__UnderlineStyle = QTextCharFormat__UnderlineStyle(4)
	QTextCharFormat__DashDotDotLine      QTextCharFormat__UnderlineStyle = QTextCharFormat__UnderlineStyle(5)
	QTextCharFormat__WaveUnderline       QTextCharFormat__UnderlineStyle = QTextCharFormat__UnderlineStyle(6)
	QTextCharFormat__SpellCheckUnderline QTextCharFormat__UnderlineStyle = QTextCharFormat__UnderlineStyle(7)
)

//go:generate stringer -type=QTextCharFormat__VerticalAlignment
//QTextCharFormat::VerticalAlignment
type QTextCharFormat__VerticalAlignment int64

const (
	QTextCharFormat__AlignNormal      QTextCharFormat__VerticalAlignment = QTextCharFormat__VerticalAlignment(0)
	QTextCharFormat__AlignSuperScript QTextCharFormat__VerticalAlignment = QTextCharFormat__VerticalAlignment(1)
	QTextCharFormat__AlignSubScript   QTextCharFormat__VerticalAlignment = QTextCharFormat__VerticalAlignment(2)
	QTextCharFormat__AlignMiddle      QTextCharFormat__VerticalAlignment = QTextCharFormat__VerticalAlignment(3)
	QTextCharFormat__AlignTop         QTextCharFormat__VerticalAlignment = QTextCharFormat__VerticalAlignment(4)
	QTextCharFormat__AlignBottom      QTextCharFormat__VerticalAlignment = QTextCharFormat__VerticalAlignment(5)
	QTextCharFormat__AlignBaseline    QTextCharFormat__VerticalAlignment = QTextCharFormat__VerticalAlignment(6)
)

//go:generate stringer -type=QTextCursor__MoveMode
//QTextCursor::MoveMode
type QTextCursor__MoveMode int64

const (
	QTextCursor__MoveAnchor QTextCursor__MoveMode = QTextCursor__MoveMode(0)
	QTextCursor__KeepAnchor QTextCursor__MoveMode = QTextCursor__MoveMode(1)
)

//go:generate stringer -type=QTextCursor__MoveOperation
//QTextCursor::MoveOperation
type QTextCursor__MoveOperation int64

const (
	QTextCursor__NoMove            QTextCursor__MoveOperation = QTextCursor__MoveOperation(0)
	QTextCursor__Start             QTextCursor__MoveOperation = QTextCursor__MoveOperation(1)
	QTextCursor__Up                QTextCursor__MoveOperation = QTextCursor__MoveOperation(2)
	QTextCursor__StartOfLine       QTextCursor__MoveOperation = QTextCursor__MoveOperation(3)
	QTextCursor__StartOfBlock      QTextCursor__MoveOperation = QTextCursor__MoveOperation(4)
	QTextCursor__StartOfWord       QTextCursor__MoveOperation = QTextCursor__MoveOperation(5)
	QTextCursor__PreviousBlock     QTextCursor__MoveOperation = QTextCursor__MoveOperation(6)
	QTextCursor__PreviousCharacter QTextCursor__MoveOperation = QTextCursor__MoveOperation(7)
	QTextCursor__PreviousWord      QTextCursor__MoveOperation = QTextCursor__MoveOperation(8)
	QTextCursor__Left              QTextCursor__MoveOperation = QTextCursor__MoveOperation(9)
	QTextCursor__WordLeft          QTextCursor__MoveOperation = QTextCursor__MoveOperation(10)
	QTextCursor__End               QTextCursor__MoveOperation = QTextCursor__MoveOperation(11)
	QTextCursor__Down              QTextCursor__MoveOperation = QTextCursor__MoveOperation(12)
	QTextCursor__EndOfLine         QTextCursor__MoveOperation = QTextCursor__MoveOperation(13)
	QTextCursor__EndOfWord         QTextCursor__MoveOperation = QTextCursor__MoveOperation(14)
	QTextCursor__EndOfBlock        QTextCursor__MoveOperation = QTextCursor__MoveOperation(15)
	QTextCursor__NextBlock         QTextCursor__MoveOperation = QTextCursor__MoveOperation(16)
	QTextCursor__NextCharacter     QTextCursor__MoveOperation = QTextCursor__MoveOperation(17)
	QTextCursor__NextWord          QTextCursor__MoveOperation = QTextCursor__MoveOperation(18)
	QTextCursor__Right             QTextCursor__MoveOperation = QTextCursor__MoveOperation(19)
	QTextCursor__WordRight         QTextCursor__MoveOperation = QTextCursor__MoveOperation(20)
	QTextCursor__NextCell          QTextCursor__MoveOperation = QTextCursor__MoveOperation(21)
	QTextCursor__PreviousCell      QTextCursor__MoveOperation = QTextCursor__MoveOperation(22)
	QTextCursor__NextRow           QTextCursor__MoveOperation = QTextCursor__MoveOperation(23)
	QTextCursor__PreviousRow       QTextCursor__MoveOperation = QTextCursor__MoveOperation(24)
)

//go:generate stringer -type=QTextCursor__SelectionType
//QTextCursor::SelectionType
type QTextCursor__SelectionType int64

const (
	QTextCursor__WordUnderCursor  QTextCursor__SelectionType = QTextCursor__SelectionType(0)
	QTextCursor__LineUnderCursor  QTextCursor__SelectionType = QTextCursor__SelectionType(1)
	QTextCursor__BlockUnderCursor QTextCursor__SelectionType = QTextCursor__SelectionType(2)
	QTextCursor__Document         QTextCursor__SelectionType = QTextCursor__SelectionType(3)
)

//go:generate stringer -type=QTextDocument__FindFlag
//QTextDocument::FindFlag
type QTextDocument__FindFlag int64

const (
	QTextDocument__FindBackward        QTextDocument__FindFlag = QTextDocument__FindFlag(0x00001)
	QTextDocument__FindCaseSensitively QTextDocument__FindFlag = QTextDocument__FindFlag(0x00002)
	QTextDocument__FindWholeWords      QTextDocument__FindFlag = QTextDocument__FindFlag(0x00004)
)

//go:generate stringer -type=QTextDocument__MetaInformation
//QTextDocument::MetaInformation
type QTextDocument__MetaInformation int64

const (
	QTextDocument__DocumentTitle QTextDocument__MetaInformation = QTextDocument__MetaInformation(0)
	QTextDocument__DocumentUrl   QTextDocument__MetaInformation = QTextDocument__MetaInformation(1)
)

//go:generate stringer -type=QTextDocument__ResourceType
//QTextDocument::ResourceType
type QTextDocument__ResourceType int64

const (
	QTextDocument__HtmlResource       QTextDocument__ResourceType = QTextDocument__ResourceType(1)
	QTextDocument__ImageResource      QTextDocument__ResourceType = QTextDocument__ResourceType(2)
	QTextDocument__StyleSheetResource QTextDocument__ResourceType = QTextDocument__ResourceType(3)
	QTextDocument__UserResource       QTextDocument__ResourceType = QTextDocument__ResourceType(100)
)

//go:generate stringer -type=QTextDocument__Stacks
//QTextDocument::Stacks
type QTextDocument__Stacks int64

const (
	QTextDocument__UndoStack         QTextDocument__Stacks = QTextDocument__Stacks(0x01)
	QTextDocument__RedoStack         QTextDocument__Stacks = QTextDocument__Stacks(0x02)
	QTextDocument__UndoAndRedoStacks QTextDocument__Stacks = QTextDocument__Stacks(QTextDocument__UndoStack | QTextDocument__RedoStack)
)

//go:generate stringer -type=QTextFormat__FormatType
//QTextFormat::FormatType
type QTextFormat__FormatType int64

const (
	QTextFormat__InvalidFormat QTextFormat__FormatType = QTextFormat__FormatType(-1)
	QTextFormat__BlockFormat   QTextFormat__FormatType = QTextFormat__FormatType(1)
	QTextFormat__CharFormat    QTextFormat__FormatType = QTextFormat__FormatType(2)
	QTextFormat__ListFormat    QTextFormat__FormatType = QTextFormat__FormatType(3)
	QTextFormat__TableFormat   QTextFormat__FormatType = QTextFormat__FormatType(4)
	QTextFormat__FrameFormat   QTextFormat__FormatType = QTextFormat__FormatType(5)
	QTextFormat__UserFormat    QTextFormat__FormatType = QTextFormat__FormatType(100)
)

//go:generate stringer -type=QTextFormat__ObjectTypes
//QTextFormat::ObjectTypes
type QTextFormat__ObjectTypes int64

const (
	QTextFormat__NoObject        QTextFormat__ObjectTypes = QTextFormat__ObjectTypes(0)
	QTextFormat__ImageObject     QTextFormat__ObjectTypes = QTextFormat__ObjectTypes(1)
	QTextFormat__TableObject     QTextFormat__ObjectTypes = QTextFormat__ObjectTypes(2)
	QTextFormat__TableCellObject QTextFormat__ObjectTypes = QTextFormat__ObjectTypes(3)
	QTextFormat__UserObject      QTextFormat__ObjectTypes = QTextFormat__ObjectTypes(0x1000)
)

//go:generate stringer -type=QTextFormat__PageBreakFlag
//QTextFormat::PageBreakFlag
type QTextFormat__PageBreakFlag int64

const (
	QTextFormat__PageBreak_Auto         QTextFormat__PageBreakFlag = QTextFormat__PageBreakFlag(0)
	QTextFormat__PageBreak_AlwaysBefore QTextFormat__PageBreakFlag = QTextFormat__PageBreakFlag(0x001)
	QTextFormat__PageBreak_AlwaysAfter  QTextFormat__PageBreakFlag = QTextFormat__PageBreakFlag(0x010)
)

//go:generate stringer -type=QTextFormat__Property
//QTextFormat::Property
type QTextFormat__Property int64

const (
	QTextFormat__ObjectIndex                       QTextFormat__Property = QTextFormat__Property(0x0)
	QTextFormat__CssFloat                          QTextFormat__Property = QTextFormat__Property(0x0800)
	QTextFormat__LayoutDirection                   QTextFormat__Property = QTextFormat__Property(0x0801)
	QTextFormat__OutlinePen                        QTextFormat__Property = QTextFormat__Property(0x810)
	QTextFormat__BackgroundBrush                   QTextFormat__Property = QTextFormat__Property(0x820)
	QTextFormat__ForegroundBrush                   QTextFormat__Property = QTextFormat__Property(0x821)
	QTextFormat__BackgroundImageUrl                QTextFormat__Property = QTextFormat__Property(0x823)
	QTextFormat__BlockAlignment                    QTextFormat__Property = QTextFormat__Property(0x1010)
	QTextFormat__BlockTopMargin                    QTextFormat__Property = QTextFormat__Property(0x1030)
	QTextFormat__BlockBottomMargin                 QTextFormat__Property = QTextFormat__Property(0x1031)
	QTextFormat__BlockLeftMargin                   QTextFormat__Property = QTextFormat__Property(0x1032)
	QTextFormat__BlockRightMargin                  QTextFormat__Property = QTextFormat__Property(0x1033)
	QTextFormat__TextIndent                        QTextFormat__Property = QTextFormat__Property(0x1034)
	QTextFormat__TabPositions                      QTextFormat__Property = QTextFormat__Property(0x1035)
	QTextFormat__BlockIndent                       QTextFormat__Property = QTextFormat__Property(0x1040)
	QTextFormat__LineHeight                        QTextFormat__Property = QTextFormat__Property(0x1048)
	QTextFormat__LineHeightType                    QTextFormat__Property = QTextFormat__Property(0x1049)
	QTextFormat__BlockNonBreakableLines            QTextFormat__Property = QTextFormat__Property(0x1050)
	QTextFormat__BlockTrailingHorizontalRulerWidth QTextFormat__Property = QTextFormat__Property(0x1060)
	QTextFormat__FirstFontProperty                 QTextFormat__Property = QTextFormat__Property(0x1FE0)
	QTextFormat__FontCapitalization                QTextFormat__Property = QTextFormat__Property(QTextFormat__FirstFontProperty)
	QTextFormat__FontLetterSpacingType             QTextFormat__Property = QTextFormat__Property(0x2033)
	QTextFormat__FontLetterSpacing                 QTextFormat__Property = QTextFormat__Property(0x1FE1)
	QTextFormat__FontWordSpacing                   QTextFormat__Property = QTextFormat__Property(0x1FE2)
	QTextFormat__FontStretch                       QTextFormat__Property = QTextFormat__Property(0x2034)
	QTextFormat__FontStyleHint                     QTextFormat__Property = QTextFormat__Property(0x1FE3)
	QTextFormat__FontStyleStrategy                 QTextFormat__Property = QTextFormat__Property(0x1FE4)
	QTextFormat__FontKerning                       QTextFormat__Property = QTextFormat__Property(0x1FE5)
	QTextFormat__FontHintingPreference             QTextFormat__Property = QTextFormat__Property(0x1FE6)
	QTextFormat__FontFamily                        QTextFormat__Property = QTextFormat__Property(0x2000)
	QTextFormat__FontPointSize                     QTextFormat__Property = QTextFormat__Property(0x2001)
	QTextFormat__FontSizeAdjustment                QTextFormat__Property = QTextFormat__Property(0x2002)
	QTextFormat__FontSizeIncrement                 QTextFormat__Property = QTextFormat__Property(QTextFormat__FontSizeAdjustment)
	QTextFormat__FontWeight                        QTextFormat__Property = QTextFormat__Property(0x2003)
	QTextFormat__FontItalic                        QTextFormat__Property = QTextFormat__Property(0x2004)
	QTextFormat__FontUnderline                     QTextFormat__Property = QTextFormat__Property(0x2005)
	QTextFormat__FontOverline                      QTextFormat__Property = QTextFormat__Property(0x2006)
	QTextFormat__FontStrikeOut                     QTextFormat__Property = QTextFormat__Property(0x2007)
	QTextFormat__FontFixedPitch                    QTextFormat__Property = QTextFormat__Property(0x2008)
	QTextFormat__FontPixelSize                     QTextFormat__Property = QTextFormat__Property(0x2009)
	QTextFormat__LastFontProperty                  QTextFormat__Property = QTextFormat__Property(QTextFormat__FontPixelSize)
	QTextFormat__TextUnderlineColor                QTextFormat__Property = QTextFormat__Property(0x2010)
	QTextFormat__TextVerticalAlignment             QTextFormat__Property = QTextFormat__Property(0x2021)
	QTextFormat__TextOutline                       QTextFormat__Property = QTextFormat__Property(0x2022)
	QTextFormat__TextUnderlineStyle                QTextFormat__Property = QTextFormat__Property(0x2023)
	QTextFormat__TextToolTip                       QTextFormat__Property = QTextFormat__Property(0x2024)
	QTextFormat__IsAnchor                          QTextFormat__Property = QTextFormat__Property(0x2030)
	QTextFormat__AnchorHref                        QTextFormat__Property = QTextFormat__Property(0x2031)
	QTextFormat__AnchorName                        QTextFormat__Property = QTextFormat__Property(0x2032)
	QTextFormat__ObjectType                        QTextFormat__Property = QTextFormat__Property(0x2f00)
	QTextFormat__ListStyle                         QTextFormat__Property = QTextFormat__Property(0x3000)
	QTextFormat__ListIndent                        QTextFormat__Property = QTextFormat__Property(0x3001)
	QTextFormat__ListNumberPrefix                  QTextFormat__Property = QTextFormat__Property(0x3002)
	QTextFormat__ListNumberSuffix                  QTextFormat__Property = QTextFormat__Property(0x3003)
	QTextFormat__FrameBorder                       QTextFormat__Property = QTextFormat__Property(0x4000)
	QTextFormat__FrameMargin                       QTextFormat__Property = QTextFormat__Property(0x4001)
	QTextFormat__FramePadding                      QTextFormat__Property = QTextFormat__Property(0x4002)
	QTextFormat__FrameWidth                        QTextFormat__Property = QTextFormat__Property(0x4003)
	QTextFormat__FrameHeight                       QTextFormat__Property = QTextFormat__Property(0x4004)
	QTextFormat__FrameTopMargin                    QTextFormat__Property = QTextFormat__Property(0x4005)
	QTextFormat__FrameBottomMargin                 QTextFormat__Property = QTextFormat__Property(0x4006)
	QTextFormat__FrameLeftMargin                   QTextFormat__Property = QTextFormat__Property(0x4007)
	QTextFormat__FrameRightMargin                  QTextFormat__Property = QTextFormat__Property(0x4008)
	QTextFormat__FrameBorderBrush                  QTextFormat__Property = QTextFormat__Property(0x4009)
	QTextFormat__FrameBorderStyle                  QTextFormat__Property = QTextFormat__Property(0x4010)
	QTextFormat__TableColumns                      QTextFormat__Property = QTextFormat__Property(0x4100)
	QTextFormat__TableColumnWidthConstraints       QTextFormat__Property = QTextFormat__Property(0x4101)
	QTextFormat__TableCellSpacing                  QTextFormat__Property = QTextFormat__Property(0x4102)
	QTextFormat__TableCellPadding                  QTextFormat__Property = QTextFormat__Property(0x4103)
	QTextFormat__TableHeaderRowCount               QTextFormat__Property = QTextFormat__Property(0x4104)
	QTextFormat__TableCellRowSpan                  QTextFormat__Property = QTextFormat__Property(0x4810)
	QTextFormat__TableCellColumnSpan               QTextFormat__Property = QTextFormat__Property(0x4811)
	QTextFormat__TableCellTopPadding               QTextFormat__Property = QTextFormat__Property(0x4812)
	QTextFormat__TableCellBottomPadding            QTextFormat__Property = QTextFormat__Property(0x4813)
	QTextFormat__TableCellLeftPadding              QTextFormat__Property = QTextFormat__Property(0x4814)
	QTextFormat__TableCellRightPadding             QTextFormat__Property = QTextFormat__Property(0x4815)
	QTextFormat__ImageName                         QTextFormat__Property = QTextFormat__Property(0x5000)
	QTextFormat__ImageWidth                        QTextFormat__Property = QTextFormat__Property(0x5010)
	QTextFormat__ImageHeight                       QTextFormat__Property = QTextFormat__Property(0x5011)
	QTextFormat__FullWidthSelection                QTextFormat__Property = QTextFormat__Property(0x06000)
	QTextFormat__PageBreakPolicy                   QTextFormat__Property = QTextFormat__Property(0x7000)
	QTextFormat__UserProperty                      QTextFormat__Property = QTextFormat__Property(0x100000)
)

//go:generate stringer -type=QTextFrameFormat__BorderStyle
//QTextFrameFormat::BorderStyle
type QTextFrameFormat__BorderStyle int64

var (
	QTextFrameFormat__BorderStyle_None       QTextFrameFormat__BorderStyle = QTextFrameFormat__BorderStyle(0)
	QTextFrameFormat__BorderStyle_Dotted     QTextFrameFormat__BorderStyle = QTextFrameFormat__BorderStyle(1)
	QTextFrameFormat__BorderStyle_Dashed     QTextFrameFormat__BorderStyle = QTextFrameFormat__BorderStyle(2)
	QTextFrameFormat__BorderStyle_Solid      QTextFrameFormat__BorderStyle = QTextFrameFormat__BorderStyle(3)
	QTextFrameFormat__BorderStyle_Double     QTextFrameFormat__BorderStyle = QTextFrameFormat__BorderStyle(4)
	QTextFrameFormat__BorderStyle_DotDash    QTextFrameFormat__BorderStyle = QTextFrameFormat__BorderStyle(5)
	QTextFrameFormat__BorderStyle_DotDotDash QTextFrameFormat__BorderStyle = QTextFrameFormat__BorderStyle(6)
	QTextFrameFormat__BorderStyle_Groove     QTextFrameFormat__BorderStyle = QTextFrameFormat__BorderStyle(7)
	QTextFrameFormat__BorderStyle_Ridge      QTextFrameFormat__BorderStyle = QTextFrameFormat__BorderStyle(8)
	QTextFrameFormat__BorderStyle_Inset      QTextFrameFormat__BorderStyle = QTextFrameFormat__BorderStyle(9)
	QTextFrameFormat__BorderStyle_Outset     QTextFrameFormat__BorderStyle = QTextFrameFormat__BorderStyle(10)
)

//go:generate stringer -type=QTextFrameFormat__Position
//QTextFrameFormat::Position
type QTextFrameFormat__Position int64

const (
	QTextFrameFormat__InFlow     QTextFrameFormat__Position = QTextFrameFormat__Position(0)
	QTextFrameFormat__FloatLeft  QTextFrameFormat__Position = QTextFrameFormat__Position(1)
	QTextFrameFormat__FloatRight QTextFrameFormat__Position = QTextFrameFormat__Position(2)
)

//go:generate stringer -type=QTextItem__RenderFlag
//QTextItem::RenderFlag
type QTextItem__RenderFlag int64

const (
	QTextItem__RightToLeft QTextItem__RenderFlag = QTextItem__RenderFlag(0x1)
	QTextItem__Overline    QTextItem__RenderFlag = QTextItem__RenderFlag(0x10)
	QTextItem__Underline   QTextItem__RenderFlag = QTextItem__RenderFlag(0x20)
	QTextItem__StrikeOut   QTextItem__RenderFlag = QTextItem__RenderFlag(0x40)
	QTextItem__Dummy       QTextItem__RenderFlag = QTextItem__RenderFlag(0xffffffff)
)

//go:generate stringer -type=QTextLayout__CursorMode
//QTextLayout::CursorMode
type QTextLayout__CursorMode int64

const (
	QTextLayout__SkipCharacters QTextLayout__CursorMode = QTextLayout__CursorMode(0)
	QTextLayout__SkipWords      QTextLayout__CursorMode = QTextLayout__CursorMode(1)
)

//go:generate stringer -type=QTextLength__Type
//QTextLength::Type
type QTextLength__Type int64

const (
	QTextLength__VariableLength   QTextLength__Type = QTextLength__Type(0)
	QTextLength__FixedLength      QTextLength__Type = QTextLength__Type(1)
	QTextLength__PercentageLength QTextLength__Type = QTextLength__Type(2)
)

//go:generate stringer -type=QTextLine__Edge
//QTextLine::Edge
type QTextLine__Edge int64

const (
	QTextLine__Leading  QTextLine__Edge = QTextLine__Edge(0)
	QTextLine__Trailing QTextLine__Edge = QTextLine__Edge(1)
)

//go:generate stringer -type=QTextLine__CursorPosition
//QTextLine::CursorPosition
type QTextLine__CursorPosition int64

const (
	QTextLine__CursorBetweenCharacters QTextLine__CursorPosition = QTextLine__CursorPosition(0)
	QTextLine__CursorOnCharacter       QTextLine__CursorPosition = QTextLine__CursorPosition(1)
)

//go:generate stringer -type=QTextListFormat__Style
//QTextListFormat::Style
type QTextListFormat__Style int64

var (
	QTextListFormat__ListDisc           QTextListFormat__Style = QTextListFormat__Style(-1)
	QTextListFormat__ListCircle         QTextListFormat__Style = QTextListFormat__Style(-2)
	QTextListFormat__ListSquare         QTextListFormat__Style = QTextListFormat__Style(-3)
	QTextListFormat__ListDecimal        QTextListFormat__Style = QTextListFormat__Style(-4)
	QTextListFormat__ListLowerAlpha     QTextListFormat__Style = QTextListFormat__Style(-5)
	QTextListFormat__ListUpperAlpha     QTextListFormat__Style = QTextListFormat__Style(-6)
	QTextListFormat__ListLowerRoman     QTextListFormat__Style = QTextListFormat__Style(-7)
	QTextListFormat__ListUpperRoman     QTextListFormat__Style = QTextListFormat__Style(-8)
	QTextListFormat__ListStyleUndefined QTextListFormat__Style = QTextListFormat__Style(0)
)

//go:generate stringer -type=QTextOption__Flag
//QTextOption::Flag
type QTextOption__Flag int64

const (
	QTextOption__ShowTabsAndSpaces                     QTextOption__Flag = QTextOption__Flag(0x1)
	QTextOption__ShowLineAndParagraphSeparators        QTextOption__Flag = QTextOption__Flag(0x2)
	QTextOption__AddSpaceForLineAndParagraphSeparators QTextOption__Flag = QTextOption__Flag(0x4)
	QTextOption__SuppressColors                        QTextOption__Flag = QTextOption__Flag(0x8)
	QTextOption__ShowDocumentTerminator                QTextOption__Flag = QTextOption__Flag(0x10)
	QTextOption__IncludeTrailingSpaces                 QTextOption__Flag = QTextOption__Flag(0x80000000)
)

//go:generate stringer -type=QTextOption__TabType
//QTextOption::TabType
type QTextOption__TabType int64

const (
	QTextOption__LeftTab      QTextOption__TabType = QTextOption__TabType(0)
	QTextOption__RightTab     QTextOption__TabType = QTextOption__TabType(1)
	QTextOption__CenterTab    QTextOption__TabType = QTextOption__TabType(2)
	QTextOption__DelimiterTab QTextOption__TabType = QTextOption__TabType(3)
)

//go:generate stringer -type=QTextOption__WrapMode
//QTextOption::WrapMode
type QTextOption__WrapMode int64

const (
	QTextOption__NoWrap                       QTextOption__WrapMode = QTextOption__WrapMode(0)
	QTextOption__WordWrap                     QTextOption__WrapMode = QTextOption__WrapMode(1)
	QTextOption__ManualWrap                   QTextOption__WrapMode = QTextOption__WrapMode(2)
	QTextOption__WrapAnywhere                 QTextOption__WrapMode = QTextOption__WrapMode(3)
	QTextOption__WrapAtWordBoundaryOrAnywhere QTextOption__WrapMode = QTextOption__WrapMode(4)
)

//go:generate stringer -type=QTouchDevice__CapabilityFlag
//QTouchDevice::CapabilityFlag
type QTouchDevice__CapabilityFlag int64

const (
	QTouchDevice__Position           QTouchDevice__CapabilityFlag = QTouchDevice__CapabilityFlag(0x0001)
	QTouchDevice__Area               QTouchDevice__CapabilityFlag = QTouchDevice__CapabilityFlag(0x0002)
	QTouchDevice__Pressure           QTouchDevice__CapabilityFlag = QTouchDevice__CapabilityFlag(0x0004)
	QTouchDevice__Velocity           QTouchDevice__CapabilityFlag = QTouchDevice__CapabilityFlag(0x0008)
	QTouchDevice__RawPositions       QTouchDevice__CapabilityFlag = QTouchDevice__CapabilityFlag(0x0010)
	QTouchDevice__NormalizedPosition QTouchDevice__CapabilityFlag = QTouchDevice__CapabilityFlag(0x0020)
	QTouchDevice__MouseEmulation     QTouchDevice__CapabilityFlag = QTouchDevice__CapabilityFlag(0x0040)
)

//go:generate stringer -type=QTouchDevice__DeviceType
//QTouchDevice::DeviceType
type QTouchDevice__DeviceType int64

const (
	QTouchDevice__TouchScreen QTouchDevice__DeviceType = QTouchDevice__DeviceType(0)
	QTouchDevice__TouchPad    QTouchDevice__DeviceType = QTouchDevice__DeviceType(1)
)

type QTransform struct {
	ptr unsafe.Pointer
}

type QTransform_ITF interface {
	QTransform_PTR() *QTransform
}

func (ptr *QTransform) QTransform_PTR() *QTransform {
	return ptr
}

func (ptr *QTransform) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QTransform) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQTransform(ptr QTransform_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTransform_PTR().Pointer()
	}
	return nil
}

func NewQTransformFromPointer(ptr unsafe.Pointer) (n *QTransform) {
	n = new(QTransform)
	n.SetPointer(ptr)
	return
}
func (ptr *QTransform) DestroyQTransform() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QTransform__TransformationType
//QTransform::TransformationType
type QTransform__TransformationType int64

const (
	QTransform__TxNone      QTransform__TransformationType = QTransform__TransformationType(0x00)
	QTransform__TxTranslate QTransform__TransformationType = QTransform__TransformationType(0x01)
	QTransform__TxScale     QTransform__TransformationType = QTransform__TransformationType(0x02)
	QTransform__TxRotate    QTransform__TransformationType = QTransform__TransformationType(0x04)
	QTransform__TxShear     QTransform__TransformationType = QTransform__TransformationType(0x08)
	QTransform__TxProject   QTransform__TransformationType = QTransform__TransformationType(0x10)
)

func (ptr *QTransform) Translate(dx float64, dy float64) *QTransform {
	if ptr.Pointer() != nil {
		tmpValue := NewQTransformFromPointer(C.QTransform_Translate(ptr.Pointer(), C.double(dx), C.double(dy)))
		qt.SetFinalizer(tmpValue, (*QTransform).DestroyQTransform)
		return tmpValue
	}
	return nil
}

func NewQTransform() *QTransform {
	tmpValue := NewQTransformFromPointer(C.QTransform_NewQTransform())
	qt.SetFinalizer(tmpValue, (*QTransform).DestroyQTransform)
	return tmpValue
}

func NewQTransform6(other QTransform_ITF) *QTransform {
	tmpValue := NewQTransformFromPointer(C.QTransform_NewQTransform6(PointerFromQTransform(other)))
	qt.SetFinalizer(tmpValue, (*QTransform).DestroyQTransform)
	return tmpValue
}

func NewQTransform5(matrix QMatrix_ITF) *QTransform {
	tmpValue := NewQTransformFromPointer(C.QTransform_NewQTransform5(PointerFromQMatrix(matrix)))
	qt.SetFinalizer(tmpValue, (*QTransform).DestroyQTransform)
	return tmpValue
}

func NewQTransform7(other QTransform_ITF) *QTransform {
	tmpValue := NewQTransformFromPointer(C.QTransform_NewQTransform7(PointerFromQTransform(other)))
	qt.SetFinalizer(tmpValue, (*QTransform).DestroyQTransform)
	return tmpValue
}

func NewQTransform3(m11 float64, m12 float64, m13 float64, m21 float64, m22 float64, m23 float64, m31 float64, m32 float64, m33 float64) *QTransform {
	tmpValue := NewQTransformFromPointer(C.QTransform_NewQTransform3(C.double(m11), C.double(m12), C.double(m13), C.double(m21), C.double(m22), C.double(m23), C.double(m31), C.double(m32), C.double(m33)))
	qt.SetFinalizer(tmpValue, (*QTransform).DestroyQTransform)
	return tmpValue
}

func NewQTransform4(m11 float64, m12 float64, m21 float64, m22 float64, dx float64, dy float64) *QTransform {
	tmpValue := NewQTransformFromPointer(C.QTransform_NewQTransform4(C.double(m11), C.double(m12), C.double(m21), C.double(m22), C.double(dx), C.double(dy)))
	qt.SetFinalizer(tmpValue, (*QTransform).DestroyQTransform)
	return tmpValue
}

func (ptr *QTransform) Reset() {
	if ptr.Pointer() != nil {
		C.QTransform_Reset(ptr.Pointer())
	}
}

func (ptr *QTransform) Map3(l core.QLine_ITF) *core.QLine {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQLineFromPointer(C.QTransform_Map3(ptr.Pointer(), core.PointerFromQLine(l)))
		qt.SetFinalizer(tmpValue, (*core.QLine).DestroyQLine)
		return tmpValue
	}
	return nil
}

func (ptr *QTransform) Map4(line core.QLineF_ITF) *core.QLineF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQLineFFromPointer(C.QTransform_Map4(ptr.Pointer(), core.PointerFromQLineF(line)))
		qt.SetFinalizer(tmpValue, (*core.QLineF).DestroyQLineF)
		return tmpValue
	}
	return nil
}

func (ptr *QTransform) Map8(path QPainterPath_ITF) *QPainterPath {
	if ptr.Pointer() != nil {
		tmpValue := NewQPainterPathFromPointer(C.QTransform_Map8(ptr.Pointer(), PointerFromQPainterPath(path)))
		qt.SetFinalizer(tmpValue, (*QPainterPath).DestroyQPainterPath)
		return tmpValue
	}
	return nil
}

func (ptr *QTransform) Map10(point core.QPoint_ITF) *core.QPoint {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFromPointer(C.QTransform_Map10(ptr.Pointer(), core.PointerFromQPoint(point)))
		qt.SetFinalizer(tmpValue, (*core.QPoint).DestroyQPoint)
		return tmpValue
	}
	return nil
}

func (ptr *QTransform) Map2(p core.QPointF_ITF) *core.QPointF {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQPointFFromPointer(C.QTransform_Map2(ptr.Pointer(), core.PointerFromQPointF(p)))
		qt.SetFinalizer(tmpValue, (*core.QPointF).DestroyQPointF)
		return tmpValue
	}
	return nil
}

func (ptr *QTransform) Map6(polygon QPolygon_ITF) *QPolygon {
	if ptr.Pointer() != nil {
		tmpValue := NewQPolygonFromPointer(C.QTransform_Map6(ptr.Pointer(), PointerFromQPolygon(polygon)))
		qt.SetFinalizer(tmpValue, (*QPolygon).DestroyQPolygon)
		return tmpValue
	}
	return nil
}

func (ptr *QTransform) Map5(polygon QPolygonF_ITF) *QPolygonF {
	if ptr.Pointer() != nil {
		tmpValue := NewQPolygonFFromPointer(C.QTransform_Map5(ptr.Pointer(), PointerFromQPolygonF(polygon)))
		qt.SetFinalizer(tmpValue, (*QPolygonF).DestroyQPolygonF)
		return tmpValue
	}
	return nil
}

func (ptr *QTransform) Map7(region QRegion_ITF) *QRegion {
	if ptr.Pointer() != nil {
		tmpValue := NewQRegionFromPointer(C.QTransform_Map7(ptr.Pointer(), PointerFromQRegion(region)))
		qt.SetFinalizer(tmpValue, (*QRegion).DestroyQRegion)
		return tmpValue
	}
	return nil
}

func (ptr *QTransform) Type() QTransform__TransformationType {
	if ptr.Pointer() != nil {
		return QTransform__TransformationType(C.QTransform_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTransform) Map9(x int, y int, tx int, ty int) {
	if ptr.Pointer() != nil {
		C.QTransform_Map9(ptr.Pointer(), C.int(int32(x)), C.int(int32(y)), C.int(int32(tx)), C.int(int32(ty)))
	}
}

func (ptr *QTransform) Map(x float64, y float64, tx float64, ty float64) {
	if ptr.Pointer() != nil {
		C.QTransform_Map(ptr.Pointer(), C.double(x), C.double(y), C.double(tx), C.double(ty))
	}
}

//go:generate stringer -type=QValidator__State
//QValidator::State
type QValidator__State int64

const (
	QValidator__Invalid      QValidator__State = QValidator__State(0)
	QValidator__Intermediate QValidator__State = QValidator__State(1)
	QValidator__Acceptable   QValidator__State = QValidator__State(2)
)

type QVector2D struct {
	ptr unsafe.Pointer
}

type QVector2D_ITF interface {
	QVector2D_PTR() *QVector2D
}

func (ptr *QVector2D) QVector2D_PTR() *QVector2D {
	return ptr
}

func (ptr *QVector2D) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QVector2D) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQVector2D(ptr QVector2D_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVector2D_PTR().Pointer()
	}
	return nil
}

func NewQVector2DFromPointer(ptr unsafe.Pointer) (n *QVector2D) {
	n = new(QVector2D)
	n.SetPointer(ptr)
	return
}
func (ptr *QVector2D) DestroyQVector2D() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQVector2D() *QVector2D {
	tmpValue := NewQVector2DFromPointer(C.QVector2D_NewQVector2D())
	qt.SetFinalizer(tmpValue, (*QVector2D).DestroyQVector2D)
	return tmpValue
}

func NewQVector2D4(point core.QPoint_ITF) *QVector2D {
	tmpValue := NewQVector2DFromPointer(C.QVector2D_NewQVector2D4(core.PointerFromQPoint(point)))
	qt.SetFinalizer(tmpValue, (*QVector2D).DestroyQVector2D)
	return tmpValue
}

func NewQVector2D5(point core.QPointF_ITF) *QVector2D {
	tmpValue := NewQVector2DFromPointer(C.QVector2D_NewQVector2D5(core.PointerFromQPointF(point)))
	qt.SetFinalizer(tmpValue, (*QVector2D).DestroyQVector2D)
	return tmpValue
}

func NewQVector2D6(vector QVector3D_ITF) *QVector2D {
	tmpValue := NewQVector2DFromPointer(C.QVector2D_NewQVector2D6(PointerFromQVector3D(vector)))
	qt.SetFinalizer(tmpValue, (*QVector2D).DestroyQVector2D)
	return tmpValue
}

func NewQVector2D7(vector QVector4D_ITF) *QVector2D {
	tmpValue := NewQVector2DFromPointer(C.QVector2D_NewQVector2D7(PointerFromQVector4D(vector)))
	qt.SetFinalizer(tmpValue, (*QVector2D).DestroyQVector2D)
	return tmpValue
}

func NewQVector2D3(xpos float32, ypos float32) *QVector2D {
	tmpValue := NewQVector2DFromPointer(C.QVector2D_NewQVector2D3(C.float(xpos), C.float(ypos)))
	qt.SetFinalizer(tmpValue, (*QVector2D).DestroyQVector2D)
	return tmpValue
}

func (ptr *QVector2D) Length() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector2D_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVector2D) X() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector2D_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVector2D) Y() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector2D_Y(ptr.Pointer()))
	}
	return 0
}

type QVector3D struct {
	ptr unsafe.Pointer
}

type QVector3D_ITF interface {
	QVector3D_PTR() *QVector3D
}

func (ptr *QVector3D) QVector3D_PTR() *QVector3D {
	return ptr
}

func (ptr *QVector3D) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QVector3D) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQVector3D(ptr QVector3D_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVector3D_PTR().Pointer()
	}
	return nil
}

func NewQVector3DFromPointer(ptr unsafe.Pointer) (n *QVector3D) {
	n = new(QVector3D)
	n.SetPointer(ptr)
	return
}
func (ptr *QVector3D) DestroyQVector3D() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQVector3D() *QVector3D {
	tmpValue := NewQVector3DFromPointer(C.QVector3D_NewQVector3D())
	qt.SetFinalizer(tmpValue, (*QVector3D).DestroyQVector3D)
	return tmpValue
}

func NewQVector3D4(point core.QPoint_ITF) *QVector3D {
	tmpValue := NewQVector3DFromPointer(C.QVector3D_NewQVector3D4(core.PointerFromQPoint(point)))
	qt.SetFinalizer(tmpValue, (*QVector3D).DestroyQVector3D)
	return tmpValue
}

func NewQVector3D5(point core.QPointF_ITF) *QVector3D {
	tmpValue := NewQVector3DFromPointer(C.QVector3D_NewQVector3D5(core.PointerFromQPointF(point)))
	qt.SetFinalizer(tmpValue, (*QVector3D).DestroyQVector3D)
	return tmpValue
}

func NewQVector3D6(vector QVector2D_ITF) *QVector3D {
	tmpValue := NewQVector3DFromPointer(C.QVector3D_NewQVector3D6(PointerFromQVector2D(vector)))
	qt.SetFinalizer(tmpValue, (*QVector3D).DestroyQVector3D)
	return tmpValue
}

func NewQVector3D7(vector QVector2D_ITF, zpos float32) *QVector3D {
	tmpValue := NewQVector3DFromPointer(C.QVector3D_NewQVector3D7(PointerFromQVector2D(vector), C.float(zpos)))
	qt.SetFinalizer(tmpValue, (*QVector3D).DestroyQVector3D)
	return tmpValue
}

func NewQVector3D8(vector QVector4D_ITF) *QVector3D {
	tmpValue := NewQVector3DFromPointer(C.QVector3D_NewQVector3D8(PointerFromQVector4D(vector)))
	qt.SetFinalizer(tmpValue, (*QVector3D).DestroyQVector3D)
	return tmpValue
}

func NewQVector3D3(xpos float32, ypos float32, zpos float32) *QVector3D {
	tmpValue := NewQVector3DFromPointer(C.QVector3D_NewQVector3D3(C.float(xpos), C.float(ypos), C.float(zpos)))
	qt.SetFinalizer(tmpValue, (*QVector3D).DestroyQVector3D)
	return tmpValue
}

func (ptr *QVector3D) Length() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector3D_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVector3D) X() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector3D_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVector3D) Y() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector3D_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVector3D) Z() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector3D_Z(ptr.Pointer()))
	}
	return 0
}

type QVector4D struct {
	ptr unsafe.Pointer
}

type QVector4D_ITF interface {
	QVector4D_PTR() *QVector4D
}

func (ptr *QVector4D) QVector4D_PTR() *QVector4D {
	return ptr
}

func (ptr *QVector4D) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QVector4D) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQVector4D(ptr QVector4D_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVector4D_PTR().Pointer()
	}
	return nil
}

func NewQVector4DFromPointer(ptr unsafe.Pointer) (n *QVector4D) {
	n = new(QVector4D)
	n.SetPointer(ptr)
	return
}
func (ptr *QVector4D) DestroyQVector4D() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
func NewQVector4D() *QVector4D {
	tmpValue := NewQVector4DFromPointer(C.QVector4D_NewQVector4D())
	qt.SetFinalizer(tmpValue, (*QVector4D).DestroyQVector4D)
	return tmpValue
}

func NewQVector4D4(point core.QPoint_ITF) *QVector4D {
	tmpValue := NewQVector4DFromPointer(C.QVector4D_NewQVector4D4(core.PointerFromQPoint(point)))
	qt.SetFinalizer(tmpValue, (*QVector4D).DestroyQVector4D)
	return tmpValue
}

func NewQVector4D5(point core.QPointF_ITF) *QVector4D {
	tmpValue := NewQVector4DFromPointer(C.QVector4D_NewQVector4D5(core.PointerFromQPointF(point)))
	qt.SetFinalizer(tmpValue, (*QVector4D).DestroyQVector4D)
	return tmpValue
}

func NewQVector4D6(vector QVector2D_ITF) *QVector4D {
	tmpValue := NewQVector4DFromPointer(C.QVector4D_NewQVector4D6(PointerFromQVector2D(vector)))
	qt.SetFinalizer(tmpValue, (*QVector4D).DestroyQVector4D)
	return tmpValue
}

func NewQVector4D7(vector QVector2D_ITF, zpos float32, wpos float32) *QVector4D {
	tmpValue := NewQVector4DFromPointer(C.QVector4D_NewQVector4D7(PointerFromQVector2D(vector), C.float(zpos), C.float(wpos)))
	qt.SetFinalizer(tmpValue, (*QVector4D).DestroyQVector4D)
	return tmpValue
}

func NewQVector4D8(vector QVector3D_ITF) *QVector4D {
	tmpValue := NewQVector4DFromPointer(C.QVector4D_NewQVector4D8(PointerFromQVector3D(vector)))
	qt.SetFinalizer(tmpValue, (*QVector4D).DestroyQVector4D)
	return tmpValue
}

func NewQVector4D9(vector QVector3D_ITF, wpos float32) *QVector4D {
	tmpValue := NewQVector4DFromPointer(C.QVector4D_NewQVector4D9(PointerFromQVector3D(vector), C.float(wpos)))
	qt.SetFinalizer(tmpValue, (*QVector4D).DestroyQVector4D)
	return tmpValue
}

func NewQVector4D3(xpos float32, ypos float32, zpos float32, wpos float32) *QVector4D {
	tmpValue := NewQVector4DFromPointer(C.QVector4D_NewQVector4D3(C.float(xpos), C.float(ypos), C.float(zpos), C.float(wpos)))
	qt.SetFinalizer(tmpValue, (*QVector4D).DestroyQVector4D)
	return tmpValue
}

func (ptr *QVector4D) SetW(w float32) {
	if ptr.Pointer() != nil {
		C.QVector4D_SetW(ptr.Pointer(), C.float(w))
	}
}

func (ptr *QVector4D) Length() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector4D_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVector4D) W() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector4D_W(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVector4D) X() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector4D_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVector4D) Y() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector4D_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVector4D) Z() float32 {
	if ptr.Pointer() != nil {
		return float32(C.QVector4D_Z(ptr.Pointer()))
	}
	return 0
}

//go:generate stringer -type=QWindow__AncestorMode
//QWindow::AncestorMode
type QWindow__AncestorMode int64

const (
	QWindow__ExcludeTransients QWindow__AncestorMode = QWindow__AncestorMode(0)
	QWindow__IncludeTransients QWindow__AncestorMode = QWindow__AncestorMode(1)
)

//go:generate stringer -type=QWindow__Visibility
//QWindow::Visibility
type QWindow__Visibility int64

const (
	QWindow__Hidden              QWindow__Visibility = QWindow__Visibility(0)
	QWindow__AutomaticVisibility QWindow__Visibility = QWindow__Visibility(1)
	QWindow__Windowed            QWindow__Visibility = QWindow__Visibility(2)
	QWindow__Minimized           QWindow__Visibility = QWindow__Visibility(3)
	QWindow__Maximized           QWindow__Visibility = QWindow__Visibility(4)
	QWindow__FullScreen          QWindow__Visibility = QWindow__Visibility(5)
)

func init() {
}
