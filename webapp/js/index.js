var user = jQuery.parseJSON(localStorage.getItem("loginUser"))
console.log(user);

if (!user) {
    document.location = "./login.html";
}

