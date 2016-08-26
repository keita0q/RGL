// var user = jQuery.parseJSON(localStorage.getItem("loginUser"))
// console.log(user);

// if (!user) {
//     document.location = "./loginnew.html";
// };
/*
var template = $('#template');
var add = $('#add');








var userss = [ {
  "id": "hoge",
  "user_id": "hgoe@gmail.com",
  "start": "2013-09-29T17:21:12Z",
  "end": "2013-09-29T17:41:12Z",
  "comment": "develop shimane",
  "tags": [
    "shimane",
    "develop"
  ]
}, 
{
  "id": "hoge1",
  "user_id": "hgoe1@gmail.com",
  "start": "2013-09-29T17:21:13Z",
  "end": "2013-09-29T17:41:13Z",
  "comment": "develop shimane1",
  "tags": [
    "shimane1",
    "develop1"
  ]
}
    ];

    
     var himaTime= null;
     var userComment= null;
     var timeA = null;
     var timeB = null;
     var himaTime = null;
     var userTags={};
     
       
 
   users.forEach(function(spare) {
       var newtemplate = template.clone(true);
        himaStart = spare.start;
      timeA = himaStart.split('T');
      timeB = timeA['1'].match(/(.{5})/g);
      himaStart= timeB;
      
      himaEnd = spare.end;
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




)};


*/



$('#ways').click(function () {
  
document.location = "./sns.html";
});
