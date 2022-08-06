<?php
// $sales = [
//     ['Northeast', '2022-01-01', '2022-02-01', '12.54'],
//     ['Northwest', '2022-01-01', '2022-02-01', '564.33'],
//     ['Southeast', '2022-01-01', '2022-02-01', '93.26'],
//     ['Southwest', '2022-01-01', '2022-02-01', '945.21'],
//     ['All Regions', '--', '--', '1597.34'],
// ];

// $fh = fopen('php://output', 'w');

// foreach ($sales as $sales_line) {
//     if (fputcsv($fh, $sales_line) === false) {
//         die("Can't write CSV line");
//     }
// }

// fclose($fh);

$fp = fopen('./csv.txt', 'r');

print "<table>\n";
while ($csv_line = fgetcsv($fp)) {
    print '<tr>';
    for ($i = 0, $j = count($csv_line); $i < $j; $i++) {
        print '<td>' . htmlentities($csv_line[$i]) . '</td>';
    }
    print "</tr>\n";
}

print "</table>\n";