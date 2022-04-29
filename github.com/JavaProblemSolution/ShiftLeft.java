class ShiftLeft{
    String str[] = {"A", "B","C","D","E","F"};
    //main method
    public static void main(String[] ars){
        ShiftLeft sLeft = new ShiftLeft();
        sLeft.remove(0);
        sLeft.remove(1);
        String[] str = sLeft.shiftLeft(str, 3);
        for(int i=0; i< str.length() ; i++)
            System.out.print(str[i] +" ");  
        
    }

     // to remove an element at a particular index
    public void remove(int index){
        if(index>=size || index<0){
            System.out.println("No element at this index");
        }else{
            for(int i=index;i<size-1;i++){
                str[i] = str[i+1];
            }
            array[size-1]=0;
            size--;
        }
    }
    
    // to trailing elements to left-shift
    public static String[] shiftLeft(String[] arr, int n){
        String newArray = new String[shiftLeft.length];
        for(int i = 0; i < shiftLeft.length; i++){
            newArray[i] = str[(i + n) % (shiftLeft.lenght + 1)];
        }
        return newArray;
    }
}
