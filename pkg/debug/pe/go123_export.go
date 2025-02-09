// export by github.com/goplus/igop/cmd/qexp

//go:build go1.23
// +build go1.23

package pe

import (
	q "debug/pe"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "pe",
		Path: "debug/pe",
		Deps: map[string]string{
			"bytes":            "bytes",
			"compress/zlib":    "zlib",
			"debug/dwarf":      "dwarf",
			"encoding/binary":  "binary",
			"errors":           "errors",
			"fmt":              "fmt",
			"internal/saferio": "saferio",
			"io":               "io",
			"os":               "os",
			"strconv":          "strconv",
			"strings":          "strings",
			"unsafe":           "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"COFFSymbol":           reflect.TypeOf((*q.COFFSymbol)(nil)).Elem(),
			"COFFSymbolAuxFormat5": reflect.TypeOf((*q.COFFSymbolAuxFormat5)(nil)).Elem(),
			"DataDirectory":        reflect.TypeOf((*q.DataDirectory)(nil)).Elem(),
			"File":                 reflect.TypeOf((*q.File)(nil)).Elem(),
			"FileHeader":           reflect.TypeOf((*q.FileHeader)(nil)).Elem(),
			"FormatError":          reflect.TypeOf((*q.FormatError)(nil)).Elem(),
			"ImportDirectory":      reflect.TypeOf((*q.ImportDirectory)(nil)).Elem(),
			"OptionalHeader32":     reflect.TypeOf((*q.OptionalHeader32)(nil)).Elem(),
			"OptionalHeader64":     reflect.TypeOf((*q.OptionalHeader64)(nil)).Elem(),
			"Reloc":                reflect.TypeOf((*q.Reloc)(nil)).Elem(),
			"Section":              reflect.TypeOf((*q.Section)(nil)).Elem(),
			"SectionHeader":        reflect.TypeOf((*q.SectionHeader)(nil)).Elem(),
			"SectionHeader32":      reflect.TypeOf((*q.SectionHeader32)(nil)).Elem(),
			"StringTable":          reflect.TypeOf((*q.StringTable)(nil)).Elem(),
			"Symbol":               reflect.TypeOf((*q.Symbol)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewFile": reflect.ValueOf(q.NewFile),
			"Open":    reflect.ValueOf(q.Open),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"COFFSymbolSize":                                 {"untyped int", constant.MakeInt64(int64(q.COFFSymbolSize))},
			"IMAGE_COMDAT_SELECT_ANY":                        {"untyped int", constant.MakeInt64(int64(q.IMAGE_COMDAT_SELECT_ANY))},
			"IMAGE_COMDAT_SELECT_ASSOCIATIVE":                {"untyped int", constant.MakeInt64(int64(q.IMAGE_COMDAT_SELECT_ASSOCIATIVE))},
			"IMAGE_COMDAT_SELECT_EXACT_MATCH":                {"untyped int", constant.MakeInt64(int64(q.IMAGE_COMDAT_SELECT_EXACT_MATCH))},
			"IMAGE_COMDAT_SELECT_LARGEST":                    {"untyped int", constant.MakeInt64(int64(q.IMAGE_COMDAT_SELECT_LARGEST))},
			"IMAGE_COMDAT_SELECT_NODUPLICATES":               {"untyped int", constant.MakeInt64(int64(q.IMAGE_COMDAT_SELECT_NODUPLICATES))},
			"IMAGE_COMDAT_SELECT_SAME_SIZE":                  {"untyped int", constant.MakeInt64(int64(q.IMAGE_COMDAT_SELECT_SAME_SIZE))},
			"IMAGE_DIRECTORY_ENTRY_ARCHITECTURE":             {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_ARCHITECTURE))},
			"IMAGE_DIRECTORY_ENTRY_BASERELOC":                {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_BASERELOC))},
			"IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT":             {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT))},
			"IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR":           {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR))},
			"IMAGE_DIRECTORY_ENTRY_DEBUG":                    {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_DEBUG))},
			"IMAGE_DIRECTORY_ENTRY_DELAY_IMPORT":             {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_DELAY_IMPORT))},
			"IMAGE_DIRECTORY_ENTRY_EXCEPTION":                {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_EXCEPTION))},
			"IMAGE_DIRECTORY_ENTRY_EXPORT":                   {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_EXPORT))},
			"IMAGE_DIRECTORY_ENTRY_GLOBALPTR":                {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_GLOBALPTR))},
			"IMAGE_DIRECTORY_ENTRY_IAT":                      {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_IAT))},
			"IMAGE_DIRECTORY_ENTRY_IMPORT":                   {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_IMPORT))},
			"IMAGE_DIRECTORY_ENTRY_LOAD_CONFIG":              {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_LOAD_CONFIG))},
			"IMAGE_DIRECTORY_ENTRY_RESOURCE":                 {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_RESOURCE))},
			"IMAGE_DIRECTORY_ENTRY_SECURITY":                 {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_SECURITY))},
			"IMAGE_DIRECTORY_ENTRY_TLS":                      {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_TLS))},
			"IMAGE_DLLCHARACTERISTICS_APPCONTAINER":          {"untyped int", constant.MakeInt64(int64(q.IMAGE_DLLCHARACTERISTICS_APPCONTAINER))},
			"IMAGE_DLLCHARACTERISTICS_DYNAMIC_BASE":          {"untyped int", constant.MakeInt64(int64(q.IMAGE_DLLCHARACTERISTICS_DYNAMIC_BASE))},
			"IMAGE_DLLCHARACTERISTICS_FORCE_INTEGRITY":       {"untyped int", constant.MakeInt64(int64(q.IMAGE_DLLCHARACTERISTICS_FORCE_INTEGRITY))},
			"IMAGE_DLLCHARACTERISTICS_GUARD_CF":              {"untyped int", constant.MakeInt64(int64(q.IMAGE_DLLCHARACTERISTICS_GUARD_CF))},
			"IMAGE_DLLCHARACTERISTICS_HIGH_ENTROPY_VA":       {"untyped int", constant.MakeInt64(int64(q.IMAGE_DLLCHARACTERISTICS_HIGH_ENTROPY_VA))},
			"IMAGE_DLLCHARACTERISTICS_NO_BIND":               {"untyped int", constant.MakeInt64(int64(q.IMAGE_DLLCHARACTERISTICS_NO_BIND))},
			"IMAGE_DLLCHARACTERISTICS_NO_ISOLATION":          {"untyped int", constant.MakeInt64(int64(q.IMAGE_DLLCHARACTERISTICS_NO_ISOLATION))},
			"IMAGE_DLLCHARACTERISTICS_NO_SEH":                {"untyped int", constant.MakeInt64(int64(q.IMAGE_DLLCHARACTERISTICS_NO_SEH))},
			"IMAGE_DLLCHARACTERISTICS_NX_COMPAT":             {"untyped int", constant.MakeInt64(int64(q.IMAGE_DLLCHARACTERISTICS_NX_COMPAT))},
			"IMAGE_DLLCHARACTERISTICS_TERMINAL_SERVER_AWARE": {"untyped int", constant.MakeInt64(int64(q.IMAGE_DLLCHARACTERISTICS_TERMINAL_SERVER_AWARE))},
			"IMAGE_DLLCHARACTERISTICS_WDM_DRIVER":            {"untyped int", constant.MakeInt64(int64(q.IMAGE_DLLCHARACTERISTICS_WDM_DRIVER))},
			"IMAGE_FILE_32BIT_MACHINE":                       {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_32BIT_MACHINE))},
			"IMAGE_FILE_AGGRESIVE_WS_TRIM":                   {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_AGGRESIVE_WS_TRIM))},
			"IMAGE_FILE_BYTES_REVERSED_HI":                   {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_BYTES_REVERSED_HI))},
			"IMAGE_FILE_BYTES_REVERSED_LO":                   {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_BYTES_REVERSED_LO))},
			"IMAGE_FILE_DEBUG_STRIPPED":                      {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_DEBUG_STRIPPED))},
			"IMAGE_FILE_DLL":                                 {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_DLL))},
			"IMAGE_FILE_EXECUTABLE_IMAGE":                    {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_EXECUTABLE_IMAGE))},
			"IMAGE_FILE_LARGE_ADDRESS_AWARE":                 {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_LARGE_ADDRESS_AWARE))},
			"IMAGE_FILE_LINE_NUMS_STRIPPED":                  {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_LINE_NUMS_STRIPPED))},
			"IMAGE_FILE_LOCAL_SYMS_STRIPPED":                 {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_LOCAL_SYMS_STRIPPED))},
			"IMAGE_FILE_MACHINE_AM33":                        {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_AM33))},
			"IMAGE_FILE_MACHINE_AMD64":                       {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_AMD64))},
			"IMAGE_FILE_MACHINE_ARM":                         {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_ARM))},
			"IMAGE_FILE_MACHINE_ARM64":                       {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_ARM64))},
			"IMAGE_FILE_MACHINE_ARMNT":                       {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_ARMNT))},
			"IMAGE_FILE_MACHINE_EBC":                         {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_EBC))},
			"IMAGE_FILE_MACHINE_I386":                        {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_I386))},
			"IMAGE_FILE_MACHINE_IA64":                        {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_IA64))},
			"IMAGE_FILE_MACHINE_LOONGARCH32":                 {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_LOONGARCH32))},
			"IMAGE_FILE_MACHINE_LOONGARCH64":                 {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_LOONGARCH64))},
			"IMAGE_FILE_MACHINE_M32R":                        {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_M32R))},
			"IMAGE_FILE_MACHINE_MIPS16":                      {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_MIPS16))},
			"IMAGE_FILE_MACHINE_MIPSFPU":                     {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_MIPSFPU))},
			"IMAGE_FILE_MACHINE_MIPSFPU16":                   {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_MIPSFPU16))},
			"IMAGE_FILE_MACHINE_POWERPC":                     {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_POWERPC))},
			"IMAGE_FILE_MACHINE_POWERPCFP":                   {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_POWERPCFP))},
			"IMAGE_FILE_MACHINE_R4000":                       {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_R4000))},
			"IMAGE_FILE_MACHINE_RISCV128":                    {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_RISCV128))},
			"IMAGE_FILE_MACHINE_RISCV32":                     {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_RISCV32))},
			"IMAGE_FILE_MACHINE_RISCV64":                     {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_RISCV64))},
			"IMAGE_FILE_MACHINE_SH3":                         {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_SH3))},
			"IMAGE_FILE_MACHINE_SH3DSP":                      {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_SH3DSP))},
			"IMAGE_FILE_MACHINE_SH4":                         {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_SH4))},
			"IMAGE_FILE_MACHINE_SH5":                         {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_SH5))},
			"IMAGE_FILE_MACHINE_THUMB":                       {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_THUMB))},
			"IMAGE_FILE_MACHINE_UNKNOWN":                     {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_UNKNOWN))},
			"IMAGE_FILE_MACHINE_WCEMIPSV2":                   {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_WCEMIPSV2))},
			"IMAGE_FILE_NET_RUN_FROM_SWAP":                   {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_NET_RUN_FROM_SWAP))},
			"IMAGE_FILE_RELOCS_STRIPPED":                     {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_RELOCS_STRIPPED))},
			"IMAGE_FILE_REMOVABLE_RUN_FROM_SWAP":             {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_REMOVABLE_RUN_FROM_SWAP))},
			"IMAGE_FILE_SYSTEM":                              {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_SYSTEM))},
			"IMAGE_FILE_UP_SYSTEM_ONLY":                      {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_UP_SYSTEM_ONLY))},
			"IMAGE_SCN_CNT_CODE":                             {"untyped int", constant.MakeInt64(int64(q.IMAGE_SCN_CNT_CODE))},
			"IMAGE_SCN_CNT_INITIALIZED_DATA":                 {"untyped int", constant.MakeInt64(int64(q.IMAGE_SCN_CNT_INITIALIZED_DATA))},
			"IMAGE_SCN_CNT_UNINITIALIZED_DATA":               {"untyped int", constant.MakeInt64(int64(q.IMAGE_SCN_CNT_UNINITIALIZED_DATA))},
			"IMAGE_SCN_LNK_COMDAT":                           {"untyped int", constant.MakeInt64(int64(q.IMAGE_SCN_LNK_COMDAT))},
			"IMAGE_SCN_MEM_DISCARDABLE":                      {"untyped int", constant.MakeInt64(int64(q.IMAGE_SCN_MEM_DISCARDABLE))},
			"IMAGE_SCN_MEM_EXECUTE":                          {"untyped int", constant.MakeInt64(int64(q.IMAGE_SCN_MEM_EXECUTE))},
			"IMAGE_SCN_MEM_READ":                             {"untyped int", constant.MakeInt64(int64(q.IMAGE_SCN_MEM_READ))},
			"IMAGE_SCN_MEM_WRITE":                            {"untyped int", constant.MakeInt64(int64(q.IMAGE_SCN_MEM_WRITE))},
			"IMAGE_SUBSYSTEM_EFI_APPLICATION":                {"untyped int", constant.MakeInt64(int64(q.IMAGE_SUBSYSTEM_EFI_APPLICATION))},
			"IMAGE_SUBSYSTEM_EFI_BOOT_SERVICE_DRIVER":        {"untyped int", constant.MakeInt64(int64(q.IMAGE_SUBSYSTEM_EFI_BOOT_SERVICE_DRIVER))},
			"IMAGE_SUBSYSTEM_EFI_ROM":                        {"untyped int", constant.MakeInt64(int64(q.IMAGE_SUBSYSTEM_EFI_ROM))},
			"IMAGE_SUBSYSTEM_EFI_RUNTIME_DRIVER":             {"untyped int", constant.MakeInt64(int64(q.IMAGE_SUBSYSTEM_EFI_RUNTIME_DRIVER))},
			"IMAGE_SUBSYSTEM_NATIVE":                         {"untyped int", constant.MakeInt64(int64(q.IMAGE_SUBSYSTEM_NATIVE))},
			"IMAGE_SUBSYSTEM_NATIVE_WINDOWS":                 {"untyped int", constant.MakeInt64(int64(q.IMAGE_SUBSYSTEM_NATIVE_WINDOWS))},
			"IMAGE_SUBSYSTEM_OS2_CUI":                        {"untyped int", constant.MakeInt64(int64(q.IMAGE_SUBSYSTEM_OS2_CUI))},
			"IMAGE_SUBSYSTEM_POSIX_CUI":                      {"untyped int", constant.MakeInt64(int64(q.IMAGE_SUBSYSTEM_POSIX_CUI))},
			"IMAGE_SUBSYSTEM_UNKNOWN":                        {"untyped int", constant.MakeInt64(int64(q.IMAGE_SUBSYSTEM_UNKNOWN))},
			"IMAGE_SUBSYSTEM_WINDOWS_BOOT_APPLICATION":       {"untyped int", constant.MakeInt64(int64(q.IMAGE_SUBSYSTEM_WINDOWS_BOOT_APPLICATION))},
			"IMAGE_SUBSYSTEM_WINDOWS_CE_GUI":                 {"untyped int", constant.MakeInt64(int64(q.IMAGE_SUBSYSTEM_WINDOWS_CE_GUI))},
			"IMAGE_SUBSYSTEM_WINDOWS_CUI":                    {"untyped int", constant.MakeInt64(int64(q.IMAGE_SUBSYSTEM_WINDOWS_CUI))},
			"IMAGE_SUBSYSTEM_WINDOWS_GUI":                    {"untyped int", constant.MakeInt64(int64(q.IMAGE_SUBSYSTEM_WINDOWS_GUI))},
			"IMAGE_SUBSYSTEM_XBOX":                           {"untyped int", constant.MakeInt64(int64(q.IMAGE_SUBSYSTEM_XBOX))},
		},
	})
}
