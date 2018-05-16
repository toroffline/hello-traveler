<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="">
    <meta name="author" content="">
    <link rel="icon" href="bootstrap-4.0.0/favicon.ico">

    <title>Floating labels example for Bootstrap</title>

    <!-- Bootstrap core CSS -->
    <link href="bootstrap-4.0.0/dist/css/bootstrap.min.css" rel="stylesheet">

    <!-- Custom styles for this template -->
    <link href="floating-labels.css" rel="stylesheet">
  </head>

  <body>
    <form class="form-signin" method="post" action="sql.php?order=<?php echo $_GET["order"]; ?>">
      <div class="text-center mb-4">
       
        <h1 class="h3 mb-3 font-weight-normal">Payment</h1>
  
      </div>

      <div class="form-label-group">

        <h5 >Name Card</h5>
        <input type="text" id="nameCard" name="nameCard" class="form-control"  required autofocus>

      </div>

      <div class="form-label-group">

        <h5>Credit/Debit</h5>
        <input type="text" name="code" pattern="{0-9}" id="card_id" class="form-control"  required autofocus maxlength="16" minlength="16">

      </div>

      <div class="form-label-group">
        <dir>
            <h5>EXP</h5>
            <h6>MONTH</h6>
            <select class="custom-select" id="inputGroupSelect01" name="day">
                <option selected>Choose...</option>
                <option value="1">01</option>
                <option value="2">02</option>
                <option value="3">03</option>
                <option value="4">04</option>
                <option value="5">05</option>
                <option value="6">06</option>
                <option value="7">07</option>
                <option value="8">08</option>
                <option value="9">09</option>
                <option value="10">10</option>
                <option value="11">11</option>
                <option value="12">12</option>
              </select>

            <h6>YEAR</h6>

            <select class="custom-select" id="inputGroupSelect01" name="month">
                <option selected>Choose...</option>
                <option value="18">18</option>
                <option value="19">19</option>
                <option value="20">20</option>
                <option value="21">21</option>
                <option value="22">22</option>
                <option value="23">23</option>
                <option value="24">24</option>
                <option value="25">25</option>
                <option value="26">26</option>
                <option value="27">27</option>
              </select>
        </div>
        <h5>CVV</h5>
        
        <input type="text" pattern="{0-9}" id="cvv" name="cvv" class="form-control"  required autofocus maxlength="3" minlength="3">
        <br>
      </div>



      <button class="btn btn-lg btn-primary btn-block" type="submit">Confirm</button>
      <p class="mt-5 mb-3 text-muted text-center">&copy; 2017-2018</p>
    </form>
  </body>
</html>
