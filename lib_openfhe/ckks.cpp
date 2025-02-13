//==================================================================================
// BSD 2-Clause License
//
// Copyright (c) 2014-2022, NJIT, Duality Technologies Inc. and other contributors
//
// All rights reserved.
//
// Author TPOC: contact@openfhe.org
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//==================================================================================

/*
  Real number serialization in a simple context. The goal of this is to show a simple setup for real number
  serialization before progressing into the next logical step - serialization and communication across
  2 separate entities
 */

#include <iomanip>
#include <tuple>
#include <unistd.h>
#include <sstream>

#include "openfhe.h"

// header files needed for serialization
#include "ciphertext-ser.h"
#include "cryptocontext-ser.h"
#include "key/key-ser.h"
#include "scheme/ckksrns/ckksrns-ser.h"

using namespace lbcrypto;

/////////////////////////////////////////////////////////////////
// NOTE:
// If running locally, you may want to replace the "hardcoded" DATAFOLDER with
// the DATAFOLDER location below which gets the current working directory
/////////////////////////////////////////////////////////////////
// char buff[1024];
// std::string DATAFOLDER = std::string(getcwd(buff, 1024));

// Save-Load locations for keys
const std::string DATAFOLDER = "ckks_demo";
std::string ccLocation       = "/cryptocontext.txt";
std::string pubKeyLocation   = "/key_pub.txt";   // Pub key
std::string skLocation       = "/key_sec.txt";   // Secret Key
std::string multKeyLocation  = "/key_mult.txt";  // relinearization key
std::string rotKeyLocation   = "/key_rot.txt";   // automorphism / rotation key

// Save-load locations for RAW ciphertexts
std::string cipherOneLocation = "/ciphertext1.txt";
std::string cipherTwoLocation = "/ciphertext2.txt";
std::string cipherLocationP = "/ciphertext";
std::string cipherLocationL = ".txt";



// Save-load locations for evaluated ciphertexts
std::string cipherMultLocation   = "/ciphertextMult.txt";
std::string reLinetLocation   = "/ciphertextReLine.txt";
std::string cipherAddLocation    = "/ciphertextAdd.txt";
std::string cipherRotLocation    = "/ciphertextRot.txt";
std::string cipherRotNegLocation = "/ciphertextRotNegLocation.txt";
std::string clientVectorLocation = "/ciphertextVectorFromClient.txt";

extern "C" {

    const char* KeyGenCKKS(int multDepth, int scaleModSize, int batchSize) {
        
        CCParams<CryptoContextCKKSRNS> parameters;
        parameters.SetMultiplicativeDepth(multDepth);
        parameters.SetScalingModSize(scaleModSize);
        parameters.SetBatchSize(batchSize);

        CryptoContext<DCRTPoly> serverCC = GenCryptoContext(parameters);

        serverCC->Enable(PKE);
        serverCC->Enable(KEYSWITCH);
        serverCC->Enable(LEVELEDSHE);

        std::cout << "Cryptocontext generated" << std::endl;

        KeyPair<DCRTPoly> serverKP = serverCC->KeyGen();
        std::cout << "Keypair generated" << std::endl;

        serverCC->EvalMultKeyGen(serverKP.secretKey);
        std::cout << "Eval Mult Keys/ Relinearization keys have been generated" << std::endl;

        serverCC->EvalRotateKeyGen(serverKP.secretKey, {1, 2, -1, -2});
        std::cout << "Rotation keys generated" << std::endl;

        if (!Serial::SerializeToFile(DATAFOLDER + ccLocation, serverCC, SerType::BINARY)) {
            std::cerr << "Error writing serialization of the crypto context to "
                        "cryptocontext.txt"
                    << std::endl;
            std::exit(1);
        }

        std::cout << "Cryptocontext serialized" << std::endl;

        if (!Serial::SerializeToFile(DATAFOLDER + pubKeyLocation, serverKP.publicKey, SerType::BINARY)) {
            std::cerr << "Exception writing public key to pubkey.txt" << std::endl;
            std::exit(1);
        }
        std::cout << "Public key serialized" << std::endl;

        if (!Serial::SerializeToFile(DATAFOLDER + skLocation, serverKP.secretKey, SerType::BINARY)) {
            std::cerr << "Exception writing secret key to skey.txt"
                        << std::endl;
            std::exit(1);
        }    
        std::cout << "Secret key serialized" << std::endl;    

        std::ofstream multKeyFile(DATAFOLDER + multKeyLocation, std::ios::out | std::ios::binary);
        if (multKeyFile.is_open()) {
            if (!serverCC->SerializeEvalMultKey(multKeyFile, SerType::BINARY)) {
                std::cerr << "Error writing eval mult keys" << std::endl;
                std::exit(1);
            }
            std::cout << "EvalMult/ relinearization keys have been serialized" << std::endl;
            multKeyFile.close();
        }
        else {
            std::cerr << "Error serializing EvalMult keys" << std::endl;
            std::exit(1);
        }

        std::ofstream rotationKeyFile(DATAFOLDER + rotKeyLocation, std::ios::out | std::ios::binary);
        if (rotationKeyFile.is_open()) {
            if (!serverCC->SerializeEvalAutomorphismKey(rotationKeyFile, SerType::BINARY)) {
                std::cerr << "Error writing rotation keys" << std::endl;
                std::exit(1);
            }
            std::cout << "Rotation keys have been serialized" << std::endl;
        }
        else {
            std::cerr << "Error serializing Rotation keys" << std::endl;
            std::exit(1);
        }


        // 拼接字符串
        std::ostringstream oss;
        oss << DATAFOLDER + ccLocation << ";" << DATAFOLDER + pubKeyLocation << ";" << DATAFOLDER + skLocation<< ";" <<DATAFOLDER + multKeyLocation<<";"<< DATAFOLDER + rotKeyLocation;
        std::string result = oss.str();
        char* cStr = (char*)malloc(result.size() + 1);
        std::strcpy(cStr, result.c_str());
        return cStr;
    }

    const char* EncryptCKKS(const char* ccloc, std::size_t length_ccloc,
                const char* pkloc, std::size_t length_pkloc,
                const double* real_parts, std::size_t length,
                int index) {
        std::string ccloc_str(ccloc, length_ccloc);
        std::string pkloc_str(pkloc, length_pkloc);
        
        std::vector<std::complex<double>> vec;
        vec.reserve(length);

        for (std::size_t i = 0; i < length; ++i) {
            vec.emplace_back(real_parts[i], 0.0);  // 虚部为0
        }

        // 初始化cryptocontext
        CryptoContext<DCRTPoly> clientCC;
        clientCC->ClearEvalMultKeys();
        clientCC->ClearEvalAutomorphismKeys();    
        lbcrypto::CryptoContextFactory<lbcrypto::DCRTPoly>::ReleaseAllContexts();
        if (!Serial::DeserializeFromFile(ccloc_str, clientCC, SerType::BINARY)) {
            std::cerr << "I cannot read serialized data from: " << ccloc_str << std::endl;
            std::exit(1);
        }
        std::cout << "client CC deserialized"<< std::endl;  


        Plaintext serverP = clientCC->MakeCKKSPackedPlaintext(vec);

        PublicKey<DCRTPoly> publicKey;
        if (!Serial::DeserializeFromFile(pkloc_str, publicKey, SerType::BINARY)) {
            std::cerr << "Exception writing public key to pubkey.txt"
                    << std::endl;
            std::exit(1);
        }
        std::cout << "publicKey deserialized"<< std::endl;    

        auto serverC = clientCC->Encrypt(publicKey, serverP);

        std::string id = std::to_string(index);
        if (!Serial::SerializeToFile(DATAFOLDER + cipherLocationP + id + cipherLocationL, serverC, SerType::BINARY)) {
            std::cerr << " Error writing ciphertext" << std::endl;
        }
        std::cout << "ciphertext serialized"<< std::endl; 
        std::string result = DATAFOLDER + cipherLocationP + id + cipherLocationL; 
        char* cStr = (char*)malloc(result.size() + 1);
        std::strcpy(cStr, result.c_str());
        return cStr;      

    }

    const char* AddCKKS(const char* ccloc, std::size_t length_ccloc,
        const char* c1loc, std::size_t length_c1loc, const char* c2loc, std::size_t length_c2loc) {
        std::string ccloc_str(ccloc, length_ccloc);
        // std::string pkloc_str(pkloc, length_pkloc);
        std::string c1loc_str(c1loc, length_c1loc);
        std::string c2loc_str(c2loc, length_c2loc);

        // 初始化cryptocontext
        CryptoContext<DCRTPoly> clientCC;
        lbcrypto::CryptoContextFactory<lbcrypto::DCRTPoly>::ReleaseAllContexts();
        if (!Serial::DeserializeFromFile(ccloc_str, clientCC, SerType::BINARY)) {
            std::cerr << "I cannot read serialized data from: " << ccloc_str << std::endl;
            std::exit(1);
        }

        // 初始化密文
        Ciphertext<DCRTPoly> clientC1;
        Ciphertext<DCRTPoly> clientC2;
        if (!Serial::DeserializeFromFile(c1loc_str, clientC1, SerType::BINARY)) {
            std::cerr << "Cannot read serialization from " << c1loc_str << std::endl;
            std::exit(1);
        }
        std::cout << "Deserialized ciphertext1" << '\n' << std::endl;

        if (!Serial::DeserializeFromFile(c2loc_str, clientC2, SerType::BINARY)) {
            std::cerr << "Cannot read serialization from " << c2loc_str << std::endl;
            std::exit(1);
        }

        std::cout << "Deserialized ciphertext2" << '\n' << std::endl;   

        // 计算加法
        auto clientCiphertextAdd    = clientCC->EvalAdd(clientC1, clientC2);

        Serial::SerializeToFile(DATAFOLDER + cipherAddLocation, clientCiphertextAdd, SerType::BINARY);
        std::cout << "Serialized all ciphertexts from client" << '\n' << std::endl;

        std::string result = DATAFOLDER + cipherAddLocation; 
        char* cStr = (char*)malloc(result.size() + 1);
        std::strcpy(cStr, result.c_str());
        return cStr;    
    }

    const char* MulCKKS(const char* ccloc, std::size_t length_ccloc,
        const char* multkloc, std::size_t length_multkloc,
        const char* c1loc, std::size_t length_c1loc, const char* c2loc, std::size_t length_c2loc) {
        std::string ccloc_str(ccloc, length_ccloc);
        // std::string pkloc_str(pkloc, length_pkloc);
        std::string c1loc_str(c1loc, length_c1loc);
        std::string c2loc_str(c2loc, length_c2loc);
        std::string multkloc_str(multkloc, length_multkloc);
        // std::string rotkloc_str(rotkloc, length_rotkloc);

        // 初始化cryptocontext
        CryptoContext<DCRTPoly> clientCC;
        clientCC->ClearEvalMultKeys();
        // clientCC->ClearEvalAutomorphismKeys();    
        lbcrypto::CryptoContextFactory<lbcrypto::DCRTPoly>::ReleaseAllContexts();
        if (!Serial::DeserializeFromFile(ccloc_str, clientCC, SerType::BINARY)) {
            std::cerr << "I cannot read serialized data from: " << ccloc_str << std::endl;
            std::exit(1);
        }

        std::ifstream multKeyIStream(multkloc_str, std::ios::in | std::ios::binary);
        std::cout << "multKeyIStream" << '\n' << std::endl;
        if (!multKeyIStream.is_open()) {
            std::cerr << "Cannot read serialization from " <<multkloc_str<< std::endl;
            std::exit(1);
        }
        std::cout << "DeserializeEvalMultKey" << '\n' << std::endl;
        if (!clientCC->DeserializeEvalMultKey(multKeyIStream, SerType::BINARY)) {
            std::cerr << "Could not deserialize eval mult key file" << std::endl;
            std::exit(1);
        }

        // 初始化密文
        Ciphertext<DCRTPoly> clientC1;
        Ciphertext<DCRTPoly> clientC2;
        if (!Serial::DeserializeFromFile(c1loc_str, clientC1, SerType::BINARY)) {
            std::cerr << "Cannot read serialization from " << c1loc_str << std::endl;
            std::exit(1);
        }
        std::cout << "Deserialized ciphertext1" << '\n' << std::endl;

        if (!Serial::DeserializeFromFile(c2loc_str, clientC2, SerType::BINARY)) {
            std::cerr << "Cannot read serialization from " << c2loc_str << std::endl;
            std::exit(1);
        }

        std::cout << "Deserialized ciphertext2" << '\n' << std::endl; 

        // 计算乘法    
        auto clientCiphertextMult   = clientCC->EvalMult(clientC1, clientC2);

        Serial::SerializeToFile(DATAFOLDER + cipherMultLocation, clientCiphertextMult, SerType::BINARY);
        std::cout << "Serialized all ciphertexts from client" << '\n' << std::endl;
        std::string result = DATAFOLDER + cipherMultLocation; 
        char* cStr = (char*)malloc(result.size() + 1);
        std::strcpy(cStr, result.c_str());
        return cStr;                  
    }

    const char* RelinearizeCKKS(const char* ccloc, std::size_t length_ccloc,
        const char* multkloc, std::size_t length_multkloc,
        const char* cloc, std::size_t length_cloc) {
        std::string ccloc_str(ccloc, length_ccloc);
        // std::string pkloc_str(pkloc, length_pkloc);
        std::string cloc_str(cloc, length_cloc);
        std::string multkloc_str(multkloc, length_multkloc);
        // std::string rotkloc_str(rotkloc, length_rotkloc);

        // 初始化cryptocontext
        CryptoContext<DCRTPoly> clientCC;
        clientCC->ClearEvalMultKeys();
        // clientCC->ClearEvalAutomorphismKeys();    
        lbcrypto::CryptoContextFactory<lbcrypto::DCRTPoly>::ReleaseAllContexts();
        if (!Serial::DeserializeFromFile(ccloc_str, clientCC, SerType::BINARY)) {
            std::cerr << "I cannot read serialized data from: " << ccloc_str << std::endl;
            std::exit(1);
        }

        std::ifstream multKeyIStream(multkloc_str, std::ios::in | std::ios::binary);
        std::cout << "multKeyIStream" << '\n' << std::endl;
        if (!multKeyIStream.is_open()) {
            std::cerr << "Cannot read serialization from " <<multkloc_str<< std::endl;
            std::exit(1);
        }
        std::cout << "DeserializeEvalMultKey" << '\n' << std::endl;
        if (!clientCC->DeserializeEvalMultKey(multKeyIStream, SerType::BINARY)) {
            std::cerr << "Could not deserialize eval mult key file" << std::endl;
            std::exit(1);
        }

        // 初始化密文
        Ciphertext<DCRTPoly> clientC1;
        if (!Serial::DeserializeFromFile(cloc_str, clientC1, SerType::BINARY)) {
            std::cerr << "Cannot read serialization from " << cloc_str << std::endl;
            std::exit(1);
        }
        std::cout << "Deserialized ciphertext1" << '\n' << std::endl;

        // 计算乘法    
        auto re = clientCC->Relinearize(clientC1);

        Serial::SerializeToFile(DATAFOLDER + reLinetLocation, re, SerType::BINARY);
        std::cout << "Serialized all ciphertexts from client" << '\n' << std::endl;
        std::string result = DATAFOLDER + reLinetLocation; 
        char* cStr = (char*)malloc(result.size() + 1);
        std::strcpy(cStr, result.c_str());
        return cStr;                  
    }

    const char* RotCKKS(const char* ccloc, std::size_t length_ccloc, 
        const char* rotkloc, std::size_t length_rotkloc,
        const char* cloc, std::size_t length_cloc, 
        int32_t index) {
        std::string ccloc_str(ccloc, length_ccloc);
        // std::string pkloc_str(pkloc, length_pkloc);
        std::string cloc_str(cloc, length_cloc);
        // std::string c2loc_str(c2loc, length_c2loc);
        // std::string multkloc_str(multkloc, length_multkloc);
        std::string rotkloc_str(rotkloc, length_rotkloc);

        // 初始化cryptocontext
        CryptoContext<DCRTPoly> clientCC;
        clientCC->ClearEvalAutomorphismKeys();    
        lbcrypto::CryptoContextFactory<lbcrypto::DCRTPoly>::ReleaseAllContexts();
        if (!Serial::DeserializeFromFile(ccloc_str, clientCC, SerType::BINARY)) {
            std::cerr << "I cannot read serialized data from: " << ccloc_str << std::endl;
            std::exit(1);
        }

        std::cout << "Deserialized rotKey" << '\n' << std::endl;
        std::ifstream rotKeyIStream(rotkloc_str, std::ios::in | std::ios::binary);
        std::cout << "rotKeyIStream" << '\n' << std::endl;
        if (!rotKeyIStream.is_open()) {
            std::cerr << "Cannot read serialization from " <<rotkloc_str<< std::endl;
            std::exit(1);
        }
        std::cout << "DeserializeEvalAutomorphismKey" << '\n' << std::endl;
        if (!clientCC->DeserializeEvalAutomorphismKey(rotKeyIStream, SerType::BINARY)) {
            std::cerr << "Could not deserialize eval rot key file" << std::endl;
            std::exit(1);
        }    

        // 初始化密文
        Ciphertext<DCRTPoly> clientC;
        if (!Serial::DeserializeFromFile(cloc_str, clientC, SerType::BINARY)) {
            std::cerr << "Cannot read serialization from " << cloc_str << std::endl;
            std::exit(1);
        }
        std::cout << "Deserialized ciphertext" << '\n' << std::endl;

        auto clientCiphertextRot    = clientCC->EvalRotate(clientC, index);


        Serial::SerializeToFile(DATAFOLDER + cipherRotLocation, clientCiphertextRot, SerType::BINARY);
        std::cout << "Serialized all ciphertexts from client" << '\n' << std::endl;
        std::string result = DATAFOLDER + cipherRotLocation; 
        char* cStr = (char*)malloc(result.size() + 1);
        std::strcpy(cStr, result.c_str());
        return cStr;    
    }

    const char* DecryptCKKS(const char* ccloc, std::size_t length_ccloc, 
        const char* skloc, std::size_t length_skloc,
        const char* cloc, std::size_t length_cloc, 
        int vectorSize) {

        std::string ccloc_str(ccloc, length_ccloc);
        std::string cloc_str(cloc, length_cloc);
        std::string skloc_str(skloc, length_skloc);

        // 初始化cryptocontext
        CryptoContext<DCRTPoly> clientCC;

        lbcrypto::CryptoContextFactory<lbcrypto::DCRTPoly>::ReleaseAllContexts();
        if (!Serial::DeserializeFromFile(ccloc_str, clientCC, SerType::BINARY)) {
            std::cerr << "I cannot read serialized data from: " << ccloc_str << std::endl;
            std::exit(1);
        }

        PrivateKey<DCRTPoly> sk;

        if (!Serial::DeserializeFromFile(skloc, sk, SerType::BINARY)) {
        std::cerr << "Exception reading secret key from skey.txt"
                    << std::endl;
        std::exit(1);
        }

        Ciphertext<DCRTPoly> ciphertext;

        Serial::DeserializeFromFile(cloc_str, ciphertext, SerType::BINARY);
        std::cout << "Deserialized ciphertext" << '\n' << std::endl;


        Plaintext plaintext;

        clientCC->Decrypt(sk, ciphertext, &plaintext);
        plaintext->SetLength(vectorSize);

        std::ostringstream oss;
        oss << plaintext; // 使用重载的 operator<< 打印到字符串流
        std::string result =  oss.str(); // 返回字符串
        char* cStr = (char*)malloc(result.size() + 1);
        std::strcpy(cStr, result.c_str());
        return cStr;      
    }

    void FreeString(const char* str) {
        free((void*)str);
    }

}