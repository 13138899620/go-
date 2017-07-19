/**
 * Created by stoneldeng on 2016/9/9.
 */

//aes
var AesHelper={
    //SHA1数字签名
    ComputeSha1Sig:function(uid,timestamp,data,signToken){
        var str=uid+timestamp+data;
        var upperCase= str.toUpperCase();
        return CryptoJS.HmacSHA1(upperCase,signToken);
    },
    //AES128 加密
    AesEncryptStr:function(source,secretKey){
        var md5=CryptoJS.MD5(secretKey).toString();
        var iv=md5.substr(0,16)
        secretKey=CryptoJS.enc.Utf8.parse(secretKey);
        iv=CryptoJS.enc.Utf8.parse(iv);
        var encrypted=CryptoJS.AES.encrypt(source,secretKey,{
            iv:iv,
            mode:CryptoJS.mode.CBC,
            padding:CryptoJS.pad.Pkcs7
        });
        encrypted=encrypted.ciphertext.toString();
        return encrypted.toString();
    },
    //AES128 解密
    AesDecryptStr:function(encryptStr,secretKey){
        var md5=CryptoJS.MD5(secretKey).toString();
        var iv=md5.substr(0,16)
        secretKey=CryptoJS.enc.Utf8.parse(secretKey);
        iv=CryptoJS.enc.Utf8.parse(iv);
        encryptStr=CryptoJS.enc.Hex.parse(encryptStr);
        encryptStr=CryptoJS.enc.Base64.stringify(encryptStr);
        var decrypted=CryptoJS.AES.decrypt(encryptStr,secretKey,{
            iv:iv,
            mode:CryptoJS.mode.CBC,
            padding:CryptoJS.pad.Pkcs7
        });
        decrypted=decrypted.toString(CryptoJS.enc.Utf8);
        return decrypted.toString()
    }
};