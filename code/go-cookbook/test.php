<?php
$salesData = [
    ['Northeast', '2022-01-01', '2022-02-01', 12.54],
    ['Northwest', '2022-01-01', '2022-02-01', 564.33],
    ['Southeast', '2022-01-01', '2022-02-01', 93.26],
    ['Southwest', '2022-01-01', '2022-02-01', 945.21],
];

$output = fopen('php://output', 'w') or die("Can't open php://output");
$total = 0;

header('Content-Type: application/csv');
header('Content-Disposition: attachment; filename="sales.csv"');

fputcsv($output, array('Region', 'Start Date', 'End Date', 'Amount'));
foreach ($salesData as $salesLine) {
    fputcsv($output, $salesLine);
    $total += $salesLine[3];
}

fputcsv($output, array('All Regions', '--', '--', $total));
fclose($output) or die("Can't open php://output");