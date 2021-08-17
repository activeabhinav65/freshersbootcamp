<?php
function maskPhoneNumber($phoneNumber) {
    $phoneNumber=strval($phoneNumber);
    $length = strlen($phoneNumber);
    $maskedString = substr($phoneNumber, 0, 2) . str_repeat("*", $length-4) . substr($phoneNumber, $length-2, 2);
    return $maskedString;
}
echo maskPhoneNumber(9810771654);
