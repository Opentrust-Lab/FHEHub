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
  Simple example for BGVrns (integer arithmetic) with serialization. Refer to
  simple-real-numbers-serial file for an example of how to use this in a "client-server" setup
 */

#include "openfhe.h"

// header files needed for serialization
#include "ciphertext-ser.h"
#include "cryptocontext-ser.h"
#include "key/key-ser.h"
#include "scheme/bgvrns/bgvrns-ser.h"

using namespace lbcrypto;

const std::string DATAFOLDER = "bgv_demo";
std::string ccLocation       = "/cryptocontext.txt";
std::string pubKeyLocation   = "/key_pub.txt";   // Pub key
std::string skLocation       = "/key_sec.txt";   // Secret Key
std::string multKeyLocation  = "/key_mult.txt";  // relinearization key
std::string rotKeyLocation   = "/key_rot.txt";   // automorphism / rotation key

std::string cipherLocationP = "/ciphertext";
std::string cipherLocationL = ".txt";

// Save-load locations for evaluated ciphertexts
std::string cipherMultLocation   = "/ciphertextMult.txt";
std::string reLinetLocation   = "/ciphertextReLine.txt";

std::string cipherAddLocation    = "/ciphertextAdd.txt";
std::string cipherRotLocation    = "/ciphertextRot.txt";


extern "C" {

        // 2 65537
        const char* KeyGenBGV(int multiplicativeDepth, int plaintextModulus) {
        
            // Sample Program: Step 1 - Set CryptoContext
            CCParams<CryptoContextBGVRNS> parameters;
            parameters.SetMultiplicativeDepth(multiplicativeDepth);
            parameters.SetPlaintextModulus(plaintextModulus);

            CryptoContext<DCRTPoly> cryptoContext = GenCryptoContext(parameters);

            // Enable features that you wish to use
            cryptoContext->Enable(PKE);
            cryptoContext->Enable(KEYSWITCH);
            cryptoContext->Enable(LEVELEDSHE);

            std::cout << "\nThe cryptocontext has been generated." << std::endl;
            // Serialize cryptocontext
            if (!Serial::SerializeToFile(DATAFOLDER + ccLocation, cryptoContext, SerType::BINARY)) {
                std::cerr << "Error writing serialization of the crypto context to "
                            "cryptocontext.txt"
                        << std::endl;
                
            }
            std::cout << "The cryptocontext has been serialized." << std::endl;

            // Sample Program: Step 2 - Key Generation

            // Initialize Public Key Containers
            KeyPair<DCRTPoly> keyPair;

            // Generate a public/private key pair
            keyPair = cryptoContext->KeyGen();

            std::cout << "The key pair has been generated." << std::endl;

            // Serialize the public key
            if (!Serial::SerializeToFile(DATAFOLDER + pubKeyLocation, keyPair.publicKey, SerType::BINARY)) {
                std::cerr << "Error writing serialization of public key to key-public.txt" << std::endl;
                
            }
            std::cout << "The public key has been serialized." << std::endl;

            // Serialize the secret key
            if (!Serial::SerializeToFile(DATAFOLDER + skLocation, keyPair.secretKey, SerType::BINARY)) {
                std::cerr << "Error writing serialization of private key to key-private.txt" << std::endl;
                
            }
            std::cout << "The secret key has been serialized." << std::endl;    

            // Generate the relinearization key
            cryptoContext->EvalMultKeyGen(keyPair.secretKey);

            std::cout << "The eval mult keys have been generated." << std::endl;

            // Serialize the relinearization (evaluation) key for homomorphic
            // multiplication
            std::ofstream emkeyfile(DATAFOLDER + multKeyLocation, std::ios::out | std::ios::binary);
            if (emkeyfile.is_open()) {
                if (cryptoContext->SerializeEvalMultKey(emkeyfile, SerType::BINARY) == false) {
                    std::cerr << "Error writing serialization of the eval mult keys to "
                                "key-eval-mult.txt"
                            << std::endl;
                    
                }
                std::cout << "The eval mult keys have been serialized." << std::endl;

                emkeyfile.close();
            }
            else {
                std::cerr << "Error serializing eval mult keys" << std::endl;
                
            }

            // Generate the rotation evaluation keys
            cryptoContext->EvalRotateKeyGen(keyPair.secretKey, {1, 2, -1, -2});

            std::cout << "The rotation keys have been generated." << std::endl;

            // Serialize the rotation keyhs
            std::ofstream erkeyfile(DATAFOLDER + rotKeyLocation, std::ios::out | std::ios::binary);
            if (erkeyfile.is_open()) {
                if (cryptoContext->SerializeEvalAutomorphismKey(erkeyfile, SerType::BINARY) == false) {
                    std::cerr << "Error writing serialization of the eval rotation keys to "
                                "key-eval-rot.txt"
                            << std::endl;
                    
                }
                std::cout << "The eval rotation keys have been serialized." << std::endl;

                erkeyfile.close();
            }
            else {
                std::cerr << "Error serializing eval rotation keys" << std::endl;
            }

            // 拼接字符串
            std::ostringstream oss;
            oss << DATAFOLDER + ccLocation << ";" << DATAFOLDER + pubKeyLocation << ";" << DATAFOLDER + skLocation<< ";" <<DATAFOLDER + multKeyLocation<<";"<< DATAFOLDER + rotKeyLocation;
            std::string result = oss.str();
            char* cStr = (char*)malloc(result.size() + 1);
            std::strcpy(cStr, result.c_str());
            return cStr;
        }

    const char* EncryptBGV(const char* ccloc, std::size_t length_ccloc,
                const char* pkloc, std::size_t length_pkloc,
                const int64_t* data, std::size_t length,
                int index) {
        std::string ccloc_str(ccloc, length_ccloc);
        std::string pkloc_str(pkloc, length_pkloc);
        
        std::vector<int64_t> vec = std::vector<int64_t>(data, data + length);

        // 初始化cryptocontext
        CryptoContext<DCRTPoly> CC;
        CC->ClearEvalMultKeys();
        CC->ClearEvalAutomorphismKeys();    
        lbcrypto::CryptoContextFactory<lbcrypto::DCRTPoly>::ReleaseAllContexts();
        if (!Serial::DeserializeFromFile(ccloc_str, CC, SerType::BINARY)) {
            std::cerr << "I cannot read serialized data from: " << ccloc_str << std::endl;
            std::exit(1);
        }
        std::cout << "CC deserialized"<< std::endl;  

        Plaintext serverP = CC->MakePackedPlaintext(vec);

        PublicKey<DCRTPoly> publicKey;
        if (!Serial::DeserializeFromFile(pkloc_str, publicKey, SerType::BINARY)) {
            std::cerr << "Exception writing public key to pubkey.txt"
                    << std::endl;
            std::exit(1);
        }
        std::cout << "publicKey deserialized"<< std::endl;    

        auto serverC = CC->Encrypt(publicKey, serverP);

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

    const char* AddBGV(const char* ccloc, std::size_t length_ccloc,
        const char* c1loc, std::size_t length_c1loc, const char* c2loc, std::size_t length_c2loc) {
        std::string ccloc_str(ccloc, length_ccloc);
        std::string c1loc_str(c1loc, length_c1loc);
        std::string c2loc_str(c2loc, length_c2loc);

        // 初始化cryptocontext
        CryptoContext<DCRTPoly> CC;
        lbcrypto::CryptoContextFactory<lbcrypto::DCRTPoly>::ReleaseAllContexts();
        if (!Serial::DeserializeFromFile(ccloc_str, CC, SerType::BINARY)) {
            std::cerr << "I cannot read serialized data from: " << ccloc_str << std::endl;
            std::exit(1);
        }

        // 初始化密文
        Ciphertext<DCRTPoly> C1;
        Ciphertext<DCRTPoly> C2;
        if (!Serial::DeserializeFromFile(c1loc_str, C1, SerType::BINARY)) {
            std::cerr << "Cannot read serialization from " << c1loc_str << std::endl;
            std::exit(1);
        }
        std::cout << "Deserialized ciphertext1" << '\n' << std::endl;

        if (!Serial::DeserializeFromFile(c2loc_str, C2, SerType::BINARY)) {
            std::cerr << "Cannot read serialization from " << c2loc_str << std::endl;
            std::exit(1);
        }

        std::cout << "Deserialized ciphertext2" << '\n' << std::endl;   

        // 计算加法
        auto CiphertextAdd    = CC->EvalAdd(C1, C2);

        Serial::SerializeToFile(DATAFOLDER + cipherAddLocation, CiphertextAdd, SerType::BINARY);
        std::cout << "Serialized CiphertextAdd" << '\n' << std::endl;

        std::string result = DATAFOLDER + cipherAddLocation; 
        char* cStr = (char*)malloc(result.size() + 1);
        std::strcpy(cStr, result.c_str());
        return cStr;    
    }

    const char* MulBGV(const char* ccloc, std::size_t length_ccloc,
        const char* multkloc, std::size_t length_multkloc,
        const char* c1loc, std::size_t length_c1loc, const char* c2loc, std::size_t length_c2loc) {
        std::string ccloc_str(ccloc, length_ccloc);
        // std::string pkloc_str(pkloc, length_pkloc);
        std::string c1loc_str(c1loc, length_c1loc);
        std::string c2loc_str(c2loc, length_c2loc);
        std::string multkloc_str(multkloc, length_multkloc);
        // std::string rotkloc_str(rotkloc, length_rotkloc);

        // 初始化cryptocontext
        CryptoContext<DCRTPoly> CC;
        CC->ClearEvalMultKeys();

        lbcrypto::CryptoContextFactory<lbcrypto::DCRTPoly>::ReleaseAllContexts();
        if (!Serial::DeserializeFromFile(ccloc_str, CC, SerType::BINARY)) {
            std::cerr << "I cannot read serialized data from: " << ccloc_str << std::endl;
            std::exit(1);
        }

        std::ifstream multKeyIStream(multkloc_str, std::ios::in | std::ios::binary);
        std::cout << "multKeyIStream" << '\n' << std::endl;
        if (!multKeyIStream.is_open()) {
            std::cerr << "Cannot read serialization multKey" <<multkloc_str<< std::endl;
            std::exit(1);
        }
        std::cout << "DeserializeEvalMultKey" << '\n' << std::endl;
        if (!CC->DeserializeEvalMultKey(multKeyIStream, SerType::BINARY)) {
            std::cerr << "Could not deserialize eval mult key file" << std::endl;
            std::exit(1);
        }

        // 初始化密文
        Ciphertext<DCRTPoly> C1;
        Ciphertext<DCRTPoly> C2;
        if (!Serial::DeserializeFromFile(c1loc_str, C1, SerType::BINARY)) {
            std::cerr << "Cannot read serialization from  " << c1loc_str << std::endl;
            std::exit(1);
        }
        std::cout << "Deserialized ciphertext1" << '\n' << std::endl;

        if (!Serial::DeserializeFromFile(c2loc_str, C2, SerType::BINARY)) {
            std::cerr << "Cannot read serialization from " << c2loc_str << std::endl;
            std::exit(1);
        }

        std::cout << "Deserialized ciphertext2" << '\n' << std::endl; 

        // 计算乘法    
        auto CiphertextMult   = CC->EvalMult(C1, C2);

        Serial::SerializeToFile(DATAFOLDER + cipherMultLocation, CiphertextMult, SerType::BINARY);
        std::cout << "Serialized all ciphertexts from " << '\n' << std::endl;
        std::string result = DATAFOLDER + cipherMultLocation; 
        char* cStr = (char*)malloc(result.size() + 1);
        std::strcpy(cStr, result.c_str());
        return cStr;                  
    }

 const char* RelinearizeBGV(const char* ccloc, std::size_t length_ccloc,
        const char* multkloc, std::size_t length_multkloc,
        const char* cloc, std::size_t length_cloc) {
        std::string ccloc_str(ccloc, length_ccloc);
        // std::string pkloc_str(pkloc, length_pkloc);
        std::string cloc_str(cloc, length_cloc);
        std::string multkloc_str(multkloc, length_multkloc);
        // std::string rotkloc_str(rotkloc, length_rotkloc);
        // 初始化cryptocontext
        CryptoContext<DCRTPoly> CC;
        CC->ClearEvalMultKeys();

        lbcrypto::CryptoContextFactory<lbcrypto::DCRTPoly>::ReleaseAllContexts();
        if (!Serial::DeserializeFromFile(ccloc_str, CC, SerType::BINARY)) {
            std::cerr << "I cannot read serialized data from: " << ccloc_str << std::endl;
            std::exit(1);
        }

        std::ifstream multKeyIStream(multkloc_str, std::ios::in | std::ios::binary);
        std::cout << "multKeyIStream" << '\n' << std::endl;
        if (!multKeyIStream.is_open()) {
            std::cerr << "Cannot read serialization multKey" <<multkloc_str<< std::endl;
            std::exit(1);
        }
        std::cout << "DeserializeEvalMultKey" << '\n' << std::endl;
        if (!CC->DeserializeEvalMultKey(multKeyIStream, SerType::BINARY)) {
            std::cerr << "Could not deserialize eval mult key file" << std::endl;
            std::exit(1);
        }

        // 初始化密文
        Ciphertext<DCRTPoly> C1;
        if (!Serial::DeserializeFromFile(cloc_str, C1, SerType::BINARY)) {
            std::cerr << "Cannot read serialization from  " << cloc_str << std::endl;
            std::exit(1);
        }
        std::cout << "Deserialized ciphertext1" << '\n' << std::endl;

        // 计算乘法    
        auto CiphertextMult   = CC->Relinearize(C1);

        Serial::SerializeToFile(DATAFOLDER + reLinetLocation, CiphertextMult, SerType::BINARY);
        std::cout << "Serialized all ciphertexts from " << '\n' << std::endl;
        std::string result = DATAFOLDER + reLinetLocation; 
        char* cStr = (char*)malloc(result.size() + 1);
        std::strcpy(cStr, result.c_str());
        return cStr;                  
    }

    const char* RotBGV(const char* ccloc, std::size_t length_ccloc, 
        const char* rotkloc, std::size_t length_rotkloc,
        const char* cloc, std::size_t length_cloc, 
        int32_t index) {
        std::string ccloc_str(ccloc, length_ccloc);
        std::string cloc_str(cloc, length_cloc);
        std::string rotkloc_str(rotkloc, length_rotkloc);

        // 初始化cryptocontext
        CryptoContext<DCRTPoly> CC;
        CC->ClearEvalAutomorphismKeys();    
        lbcrypto::CryptoContextFactory<lbcrypto::DCRTPoly>::ReleaseAllContexts();
        if (!Serial::DeserializeFromFile(ccloc_str, CC, SerType::BINARY)) {
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
        if (!CC->DeserializeEvalAutomorphismKey(rotKeyIStream, SerType::BINARY)) {
            std::cerr << "Could not deserialize eval rot key file" << std::endl;
            std::exit(1);
        }    

        // 初始化密文
        Ciphertext<DCRTPoly> C;
        if (!Serial::DeserializeFromFile(cloc_str, C, SerType::BINARY)) {
            std::cerr << "Cannot read serialization from " << cloc_str << std::endl;
            std::exit(1);
        }
        std::cout << "Deserialized ciphertext" << '\n' << std::endl;

        auto CiphertextRot    = CC->EvalRotate(C, index);


        Serial::SerializeToFile(DATAFOLDER + cipherRotLocation, CiphertextRot, SerType::BINARY);
        std::cout << "Serialized all ciphertexts from " << '\n' << std::endl;
        std::string result = DATAFOLDER + cipherRotLocation; 
        char* cStr = (char*)malloc(result.size() + 1);
        std::strcpy(cStr, result.c_str());
        return cStr;    
    }

    const char* DecryptBGV(const char* ccloc, std::size_t length_ccloc, 
        const char* skloc, std::size_t length_skloc,
        const char* cloc, std::size_t length_cloc, 
        int vectorSize) {

        std::string ccloc_str(ccloc, length_ccloc);
        std::string cloc_str(cloc, length_cloc);
        std::string skloc_str(skloc, length_skloc);

        // 初始化cryptocontext
        CryptoContext<DCRTPoly> CC;

        lbcrypto::CryptoContextFactory<lbcrypto::DCRTPoly>::ReleaseAllContexts();
        if (!Serial::DeserializeFromFile(ccloc_str, CC, SerType::BINARY)) {
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

        CC->Decrypt(sk, ciphertext, &plaintext);
        plaintext->SetLength(vectorSize);

        std::ostringstream oss;
        oss << plaintext; // 使用重载的 operator<< 打印到字符串流
        std::string result =  oss.str(); // 返回字符串
        char* cStr = (char*)malloc(result.size() + 1);
        std::strcpy(cStr, result.c_str());
        return cStr;      
    }
}