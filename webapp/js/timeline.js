var user = jQuery.parseJSON(localStorage.getItem("loginUser"))
console.log(user);

if (!user) {
    document.location = "./loginnew.html";
};

var template = $('#template');
var add = $('#add');

var now = new Date();
var year = now.getFullYear();
var month = now.getMonth()+1;
var day = now.getDate();
console.log(day);
var hour = now.getHours();
var min = now.getMinutes();
var sec = now.getSeconds();
//profile.start = "2016-08-26T"+"17:21:12"+"Z";
var spares = [];
var profile = {time:sprintf("%4d-%02d-%02dT%02d:%02d:%02dZ", year,month,day,hour,min,sec)};
    console.log(profile);
    
window.Api.ajaxSearchSpares(profile, function(response) {
    console.log(response);
    
     var himaTime= null;
     var userComment= null;
     var timeA = null;
     var timeB = null;
     var himaTime = null;
     var userTags={};
     
   response.forEach(function(spare) {
       var newtemplate = template.clone(true);
        var himaStart = spare.start;
      timeA = himaStart.split('T');
      timeB = timeA['1'].match(/(.{5})/g);
      himaStart= timeB;
      
      var himaEnd = spare.end;
      timeA = himaEnd.split('T');
      timeB = timeA['1'].match(/(.{5})/g);
      himaEnd= timeB;
      
      himaTime = 'Spare Time : ' + himaStart + ' - ' + himaEnd;
      
         userComment = 'Comment : ' + spare.comment;
         userTags = 'Tags : ' +  spare.tags;
        
        newtemplate.find('p').get(0).innerHTML = himaTime;
        newtemplate.find('p').get(1).innerHTML = userComment;
        newtemplate.find('p').get(2).innerHTML = userTags;
        newtemplate.css('display', 'block');
    add.append(newtemplate);
    
    
});
});

$('#ways').click(function () {
document.location = "./sns.html";
});

// searchボタン
$('#searchIcon').click(function () {
    console.log("click")
 var searchTag = $('#searchBox').val();
 var p = {};
 console.log(searchTag);
 p.tag = searchTag;
 window.Api.ajaxSearchSpares(p,function(res){
    console.log(res);
    
     var himaTime= null;
     var userComment= null;
     var timeA = null;
     var timeB = null;
     var himaTime = null;
     var userTags={};
     
   res.forEach(function(spare) {
       var newtemplate = template.clone(true);
        var himaStart = spare.start;
      timeA = himaStart.split('T');
      timeB = timeA['1'].match(/(.{5})/g);
      himaStart= timeB;
      
      var himaEnd = spare.end;
      timeA = himaEnd.split('T');
      timeB = timeA['1'].match(/(.{5})/g);
      himaEnd= timeB;
      
      himaTime = 'Spare Time : ' + himaStart + ' - ' + himaEnd;
      
         userComment = 'Comment : ' + spare.comment;
         userTags = 'Tags : ' +  spare.tags;
        
        newtemplate.find('p').get(0).innerHTML = himaTime;
        newtemplate.find('p').get(1).innerHTML = userComment;
        newtemplate.find('p').get(2).innerHTML = userTags;
        newtemplate.css('display', 'block');
    add.append(newtemplate);
   }); 
 });
});