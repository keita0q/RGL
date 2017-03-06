//新規投稿画面
 


//データ取得
$('#submitc').click(function () {
    // console.log("come");
     var profile = {};
    //var id = $('#spareId2').val();
    // var user_id = $('#spareuser_id').val();
     var start = $('#start2').val();
    // var end = $('#end').val();
    // var comment = $('#comment').val();

    //   profile.id = "hogehoge";
    //   profile.user_id = "hgoe@gmail.com";
     //  profile.time = "2013-09-29T"+"17:21:13"+"Z";
    //   profile.end = "2013-09-29T"+"17:41:12"+"Z";
    //   profile.Comment = "develop shimane";
       profile.tag = "shimane";
      
    //   profile.id = id;
    //   profile.user_id = user_id;
      // profile.start = "2016-08-26T"+start+":00Z";
    //   profile.end = "2016-08-26T"+end+":00Z";
    //   profile.Comment = comment;
    //   profile.tags = ["shimane","develop"];
    
    
  
      
    //console.log(profile);

    window.Api.ajaxSearchSpares(profile,function(response) {
            //$('#hoge').html(response);
            //$('#result').html(response);
            //alert(JSON.stringify(response));
        });
    //console.log('ajax');
    return false;

   });