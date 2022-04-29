import java.util.Arrays;
public class SetDynamicArray{
    private int array[];
    private int size;
    private int capacity;
     
    //constructor
    public SetDynamicArray(){
        array = new int[2];
        size=0;
        capacity=2;
    }

     // main method
     public static void main(String[] args){
        DynamicArray dArray = new DynamicArray();
        dArray.set(0, 10);
        dArray.set(1, 20);
       
    }

       
    // to add an element at a particular index
    public void set(int index, int element){
        // double the capacity if all the allocated space is utilized
        if (size == capacity){
            ensureCapacity(2); 
        }
        // shift all elements from the given index to right
        for(int i=size-1;i>=index;i--){
            array[i+1] = array[i];
        }
        // insert the element at the specified index
        array[index] = element;
        size++;
    }
 
    
    /* method to increase the capacity, if necessary, to ensure it can hold at least the 
    *  number of elements specified by minimum capacity arguement
    */
    public void ensureCapacity(int minCapacity){
        int temp[] = new int[capacity*minCapacity];
        for (int i=0; i < capacity; i++){
            temp[i] = array[i];
        }
        array = temp;
        capacity = capacity * minCapacity;
    }
     
   
}
