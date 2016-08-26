//新規登録画面
// 渡す値 name, age, userId, userPas, line, tw, fa, mail, tags
// id= #name, #age, #userId, #userPas, #sns, #tags ,#submit

// すべて押されていたらsubmit表示

$(function() {
      //snsの追加
   
  

$('#lineT').prop('disabled', true);
$('#twT').prop('disabled', true);
$('#faT').prop('disabled', true);
$('#mailT').prop('disabled', true);
$('#line').click(function() {
        if ($(this).prop('checked') == false) {

            $('#lineT').prop('disabled', true);

        } else {
           
            $('#lineT').prop('disabled', false);
        }
    });

$('#tw').click(function() {
        if ($(this).prop('checked') == false) {

            $('#twT').prop('disabled', true);

        } else {
           
            $('#twT').prop('disabled', false);
        }
    });

$('#fa').click(function() {
        if ($(this).prop('checked') == false) {

            $('#faT').prop('disabled', true);

        } else {
           
            $('#faT').prop('disabled', false);
        }
    });
$('#mail').click(function() {
        if ($(this).prop('checked') == false) {

            $('#mailT').prop('disabled', true);

        } else {
           
            $('#mailT').prop('disabled', false);
        }
    });
    
    $('#submita').prop('disabled', true);
   
    $('#name,#userId,#userPas,#LineT,#twT,#faT,#mailT,#tags').on('keydown keyup keypress change', function() {
     
         var sns = $('#lineT').val().length > 0 || $('#twT').val().length > 0 || $('#faT').val().length > 0 || $('#mailT').val().length > 0;

        if ($('#name').val().length > 0 && $('#age').val().length > 0  && $('#userId').val().length > 0 && $('#userPas').val().length > 0 && sns){
            $('#submita').prop('disabled', false);

        } else {

            $('#submita').prop('disabled', true);
        
        }
    });
});

//データ取得
$('#submita').click(function () {
    console.log("come");
    var profile = {};
    var name = $('#name').val();
    var age = $('#age').val();
    var userId = $('#userId').val();
    var userPas = $('#userPas').val();
    var sex = $('input[name=optionsRadios]:checked').val();
    profile.sex = sex;
    profile.id = userId;
      profile.password = userPas;
      profile.name = name;
      profile.age = parseInt(age);
    
    var line = null;
    var tw = null;
    var fa = null;
    var mail = null;
    
    if ($('#lineT').val().length > 0){

         line = $('#lineT').val();
         profile['sns'] = {line : line};
   
   };
   
    profile.sns = {twitter : "", facebook: "", email: "hoge", line:""}
  if ($('#twT').val().length > 0){
         tw = $(this).val();
        profile['sns'] = {twitter : tw};
   };
   
   if ($('#faT').val().length > 0){
         fa = $(this).val();
         profile['sns'] = {facebook : fa};
   };
   if ($('#mailT').val().length > 0){
         mail = $(this).val();
         profile['sns'] = {email : mail};
   };
   
    var tags = $('#tags').val();
    
    profile.tags = tags.split(',');;
  
      
    console.log(profile);
    alert('koko');
    window.Api.ajaxPostUser(profile,function(response) {
            document.location = "index.html";
        });
    return false;

   });
   

 
    


