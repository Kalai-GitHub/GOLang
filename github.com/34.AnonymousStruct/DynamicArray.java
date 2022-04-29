import java.util.Arrays;
public class DynamicArray{
    private int array[];
    private int size;
    private int capacity;
     
    public DynamicArray(){
        array = new int[2];
        size=0;
        capacity=2;
    }
     // main method
     public static void main(String[] args){
        DynamicArray dArray = new DynamicArray();
        dArray.append(10);
        dArray.append(20);
        dArray.append(30);
        dArray.append(40);
        System.out.println(dArray.get(0));
        System.out.println(dArray.get(1));
    }

    // to add an element at the end
    public void append(int value){
        // double the capacity if all the allocated space is utilized
        if (size == capacity){
            ensureCapacity(2); 
        }
        array[size] = element;
        size++;
    }
     
    // to get an element at an index
    public int get(int index){
        return array[index];
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
