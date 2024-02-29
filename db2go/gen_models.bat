@echo off


set OUT_DIR=..
set PACK_NAME="models"
set SUFFIX_NAME="do"
set READ_ONLY="created_time,updated_time"
set DSN_URL="mysql://root:123456@192.168.1.16:3306/coin-ant?charset=utf8"
set JSON_PROPERTIES="omitempty"
set SPEC_TYPES="rich_list.extra_data=ExtraData"
set TINYINT_TO_BOOL="deleted, is_ok"
set IMPORT_MODELS="meta-market/pkg/dal/models"
db2go --url %DSN_URL% --out %OUT_DIR%  --enable-decimal --spec-type %SPEC_TYPES% ^
--suffix %SUFFIX_NAME% --package %PACK_NAME% --readonly %READ_ONLY%  --tinyint-as-bool %TINYINT_TO_BOOL% ^
--dao dao --import-models %IMPORT_MODELS%


If "%errorlevel%" == "0" (
echo generate go file ok, formatting...
gofmt -w %OUT_DIR%/%PACK_NAME%
) else (
echo if there is no db2go.exe, please download from https://github.com/civet148/release/tree/master/db2go
)

pause