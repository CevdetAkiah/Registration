
//validateForm checks if the password is equal to the repeat password fields
function validateForm() {
    let x = document.forms["formReg"]["pwd"].value;
    let y = document.forms["formReg"]["pwd-repeat"].value;
    if (x !== y) {
        Swal.fire({
            icon: 'error',
            title: `Oops...`,
            text: `Your password doesn't match.`,
        })

        return false;
    }
}

//validateInput changes css styles depending on if the input is valid.
function validateInput(x) {
    let myEl = document.getElementById(x);
    if (myEl.type == "email") {
        if (!myEl.checkValidity()) {
            myEl.classList.add("invalid");
            myEl.classList.remove("valid");
        } else {
            myEl.classList.remove("invalid");
            myEl.classList.add("valid");
        }
    } else {
        if (!myEl.checkValidity()) {
            myEl.classList.add("invalid");
            myEl.classList.remove("valid");
        } else {
            myEl.classList.remove("invalid");
            myEl.classList.add("valid");
        }
    }
}




// function inputValidate() {
//     form = document.forms["formReg"];
//     field = Array.from(form.elements);

//     field.forEach(i => {
//         i.setCustomValidity("");
//         if (length(i) > 0) {
//             if (!i.checkValidity()) {
//                 i.classList.add('invalid');
//             }
//         }
//     })
// }