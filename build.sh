#!/bin/bash

# 设置编译器和共享库输出目录
CXX=g++
OUTPUT_DIR=./api
SOURCE_DIR=./lib_openfhe
INCLUDE_DIR=/usr/local/include/openfhe
LIB_DIR=/usr/local/lib
FLAGS="-shared -fPIC -Wl,-rpath=${LIB_DIR}"

# 包含路径
INCLUDE_FLAGS="-I${INCLUDE_DIR}/core -I${INCLUDE_DIR}/pke -I${INCLUDE_DIR} -I${INCLUDE_DIR}/binfhe"

# 链接的库
LIBS="-lOPENFHEpke -lOPENFHEbinfhe -lOPENFHEcore"

# 编译 bfv.cpp 并生成 libbfv.so
echo "Compiling bfv.cpp to libbfv.so..."
$CXX $FLAGS -o ${OUTPUT_DIR}/libbfv.so $INCLUDE_FLAGS -L${LIB_DIR} ${LIBS} ${SOURCE_DIR}/bfv.cpp

# 编译 bgv.cpp 并生成 libbgv.so
echo "Compiling bgv.cpp to libbgv.so..."
$CXX $FLAGS -o ${OUTPUT_DIR}/libbgv.so $INCLUDE_FLAGS -L${LIB_DIR} ${LIBS} ${SOURCE_DIR}/bgv.cpp

# 编译 ckks.cpp 并生成 libckks.so
echo "Compiling ckks.cpp to libckks.so..."
$CXX $FLAGS -o ${OUTPUT_DIR}/libckks.so $INCLUDE_FLAGS -L${LIB_DIR} ${LIBS} ${SOURCE_DIR}/ckks.cpp

# 编译 Go 项目 keyxx_server.go，使用 CGO 编译选项
echo "Building keyxx_server.go..."
CGO_CFLAGS="-I/usr/include" go build keyxx_server.go

if [ $? -eq 0 ]; then
    echo "Go build succeeded!"
else
    echo "Go build failed!"
    exit 1
fi

echo "Compilation completed!"
