set fp [open "input_1" r]
set data [read $fp]
close $fp

set previous 0

set data2 [split $data "\n"]
set found 0
set numlines 0
foreach line $data2 {
    incr numlines
    if { $previous > 0 && $line > $previous } {
        incr found
    }
    set previous $line
}

puts "Solution 1: $found"


set previous 0
set found 0

for {set i 0} {$i < [expr {$numlines-3}]} {incr i}  {
    set a [lindex $data2 $i]
    set b [lindex $data2 [expr {$i + 1}]]
    set c [lindex $data2 [expr {$i + 2}]]

    set sum [expr {$a + $b + $c}]
    if { $previous > 0 && $sum > $previous } {
        incr found
    }
    set previous $sum
}

puts "Solution 2: $found"
