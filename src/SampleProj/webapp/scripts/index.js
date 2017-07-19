$(document).ready(function () {
    //登录写COOKIE状态
    $.get("api/account/Login",{},function(success){
        //1. 获取用户票据信息TICKET，用于加解密
        var ticket={};
        $.get("/api/user/GetUserTicket",{},function(data){
            var resultModel=JSON.parse(data);
            if(resultModel.StatusCode==200){
                ticket=JSON.parse(resultModel.Data)
                //alert(resultModel.Data)
            }
        });

        //绑定按钮事件
        //获取用户信息
        $("#getuserlist").on("click", function () {
            $.get("/api/user/GetUserList",{},function (result) {
                var dataList=JSON.parse(result);
                if(dataList.StatusCode==200){
                    var html="";
                    var userList=JSON.parse( dataList.Data);
                    for(var i=0;i<userList.length;i++){
                        html+='<li>'+userList[i].UserName+'</li>'
                    }
                    $("#user_list").html(html);
                }else{
                    $("#user_list").html(dataList.Data.toString());
                }
            });
        });

        //获取用户列表，过程加密
        $("#getuserlistencrypt").on("click", function () {
            $.get("/api/user/GetEncryptedUserList",{},function (result) {
                var dataList=JSON.parse(result);
                if(dataList.StatusCode==200){
                    var html="";
                    var  userList;
                    if(dataList.IsEncrypted){//有加密的
                        //解密字符串
                        var decryptStr= AesHelper.AesDecryptStr(dataList.Data,ticket.SecretKey);
                         userList=JSON.parse(decryptStr);
                    }else{//无加密传输
                         userList=JSON.parse( dataList.Data);
                    }
                    for(var i=0;i<userList.length;i++){
                        html+='<li>'+userList[i].UserName+'</li>'
                    }
                    $("#user_list_encrypt").html(html);
                }else{
                    $("#user_list_encrypt").html(dataList.Data.toString());
                }
            });
        });

        //添加用户操作
        $("#btn_addUser").on("click",function(){
            var user={UserName:""};
            user.UserName= $.trim($("#userName").val());
            $.post("/api/user/AddUser",{item:JSON.stringify(user)},function(data){
                alert(data)
            })
        });

        //添加用户传输过程加密
        $("#btn_addUser_encrypt").on("click",function(){
            var timeStamp=new Date().getTime();
            var user={UserName:""};
            user.UserName= $.trim($("#userName_encrypt").val());
            //加密POST数据
            var encrypt=AesHelper.AesEncryptStr(JSON.stringify(user),ticket.SecretKey);
            $.ajax({
                type:"POST",
                url:"/api/user/AddEncryptedUser",
                data:{item:encrypt},
                xhrFields:{
                    withCredentials:true
                },
                headers:{
                    "uid":ticket.UId,
                    "go-timestamp":timeStamp.toString(),
                    "go-signature":AesHelper.ComputeSha1Sig(ticket.UId,timeStamp.toString(),encrypt,ticket.SigToken)
                },
                success:function(data){
                    //TODO 返回处理
                    alert(data)
                },
                error:function(jqXHR,textStatus,errorThrown){
                   //TODO 出错处理
                }
            });
        });

    });
});