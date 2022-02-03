function validateForm() {
    let x = document.forms["formReg"]["pwd"].value;
    let y = document.forms["formReg"]["pwd-repeat"].value;
    if (x !== y) {
        Swal.fire({
            icon: 'error',
            title: `Oops...`,
            text: `Your password doesn't match`,
            footer: '<a href="">Why do I have this issues?</a>'
        })
        return false;
    }
}