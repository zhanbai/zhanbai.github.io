<?php
$words = preg_split('/\d\. /', 'my day: 1. get up 2. get dressed 3. eat toast');
var_export($words);
$words = preg_split('/ x /i', '31 inches x 22 inches X 9 inches');
var_export($words);