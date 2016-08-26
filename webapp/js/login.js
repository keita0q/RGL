$('#login').click(function () {
    var userId = $('#userId').val();
    var userPas = $('#userPas').val();
    window.Api.ajaxGetUser(userId,function(res){
        console.log(res);
        localStorage.setItem("loginUser",JSON.stringify(res));
    },function(){
        document.location = "./index.html";
    });
    return false;
   });
 
 
// $('#newAccount').click(function () {
//   document.location = "loginomori.html";
// });


