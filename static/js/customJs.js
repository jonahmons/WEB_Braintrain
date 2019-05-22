$(function(){

 $('#password, #confirm_password').on('keyup', function () {
  if ($('#password').val() == $('#confirm_password').val()) {
    $('#password_message').addClass("is-invisible");
  } else {
    $('#password_message').removeClass("is-invisible");
  }
  });
});
