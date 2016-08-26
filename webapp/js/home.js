// var user = jQuery.parseJSON(localStorage.getItem("loginUser"))
// console.log(user);
// $(".userName").text(user.name);
var isDisabled = true;
$('input').prop('disabled', isDisabled);
$("#save").css("display","none");

$('#himaB').click(function(){
   document.location = "./spare_form.html"; 
   return false;
});

$("#edit").click(function(){
    //console.log("click")
    isDisabled = !isDisabled;
    $('input').prop('disabled', isDisabled);
    $("#edit").css("display","none");
    $("#save").css("display","");  
})

$("#save").click(function() {
    isDisabled = !isDisabled;
    $('input').prop('disabled', isDisabled);
    $("#save").css("display","none");
    $("#edit").css("display","");  
})