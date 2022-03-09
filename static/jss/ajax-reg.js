var form = document.querySelector('#form');
var email = document.querySelector('#email');
var p1 = document.querySelector('#pwd');
var p2 = document.querySelector('#pwd-repeat');
var btn = document.querySelector('#submitbutton');

var emailErr = document.querySelector('#email-err');
var pwdErr = document.querySelector('#pwd-err');
var pwdErr2 = document.querySelector('#pwd-rpt-err');

// checks if email is already registered
email.addEventListener('input', function () {
    console.log(email.value);
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/checkUserName', true);
    xhr.addEventListener('readystatechange', function () {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            var item = xhr.responseText;
            console.log(item);
            if (item == 'true') {
                emailErr.textContent = 'Email already registered - try log in again!';
                document.getElementById('email').classList.add("invalid");
                document.getElementById('email').classList.remove("valid");
                btn.disabled = true;
            } else {
                emailErr.textContent = '';
                btn.disabled = false;
            }
        }
    });
    xhr.send(email.value);
});
