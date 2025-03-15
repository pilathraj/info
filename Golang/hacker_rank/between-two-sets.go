// https://www.hackerrank.com/challenges/between-two-sets/problem?isFullScreen=true

// https://studyalgorithms.com/array/hackerrank-between-two-sets/

func gcd(a int32, b int32) int32{
    
    for b !=0{
        a, b = b, a%b
    }
    return a
}

func lcd(a int32, b int32) int32 {
    return a*b/gcd(a,b)
}


func getTotalX(a []int32, b []int32) int32 {
    var lcd_value, gcd_value, result int32 = a[0], b[0], 0
    for _, v := range(a){
       lcd_value =  lcd(lcd_value, v)
    }
    for _, v := range(b){
       gcd_value =  gcd(gcd_value, v)
    }
    
    var multiple int32 = 0;
    for multiple <= gcd_value {
      multiple += lcd_value;

      if gcd_value % multiple == 0 {
        result = result+1 
      }
    }

     fmt.Println(lcd_value, gcd_value, result) // 4, 16, 3
    return result;
    
   
    
}
