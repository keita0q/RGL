//新規投稿画面
var now = new Date();
var year = now.getFullYear();
var month = now.getMonth();
var day = now.getDate();
var hour = now.getHours();
var min = now.getMinutes();
var sec = now.getSeconds();
var user = jQuery.parseJSON(localStorage.getItem("loginUser"));
$('#spareId').text(user.id + now);
$('#spareuser_id').text(user.id);

 


//データ取得
$('#submitb').click(function () {
    console.log("come");
    var profile = {};
    var id = $('#spareId').text();
    var user_id = $('#spareuser_id').text();
    var start = $('#start').val();
    var end = $('#end').val();
    var comment = $('#comment').val();
    var tags =$('#tagss').val();

      // profile.id = "hogehoge";
      // profile.user_id = "hgoe@gmail.com";
      // profile.start = "2016-08-26T"+"17:21:12"+"Z";
      // profile.end = "2016-08-26T"+"17:41:12"+"Z";
      // profile.Comment = "develop shimane";
      // profile.tags = ["shimane","develop"];
      
      profile.id = id;
      profile.user_id = user_id;
      profile.start = year+"-"+("0"+(month+1)).slice(-2)+"-"+("0"+day).slice(-2)+"T"+start+":00Z";
      profile.end = year+"-"+("0"+(month+1)).slice(-2)+"-"+("0"+day).slice(-2)+"T"+end+":00Z";
      profile.Comment = comment;
      profile.tags = tags.split(',');
      
    
    
  
      
    console.log(profile);

    window.Api.ajaxSaveSpares(profile,function(response) {
            //$('#hoge').html(response);
            alert('ajax');
            document.location = "panels1.html";
        });
    return false;
    
   });
   

