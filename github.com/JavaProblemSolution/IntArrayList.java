import java.io.*;
import java.util.*;
 
class IntArrayList {
    int n = 5;
    ArrayList<Integer> arrlist = new ArrayList<Integer>(n);

    public static void main(String[] args)
    {
        // Add elements in the arrayList
        addArrList(arrli)
        //Print the elements in the arrayList
        print(arrlist)
        // Remove element using index
        removeArrList(3)
        
    }

    private void addArrList(int[] arrlist) {
        for (int i = 1; i <= n; i++)
            arrlist.add(i);
 
        System.out.println(arrlist);
    }

    private void removeArrList(int index) {
        for (int i = 1; i <= n; i++)
            if i == index{
                arrli.remove(3);

            }
 
        System.out.println(arrlist);
    }
    private void print(int[] arrlist) {
        for (int i = 0; i < arrli.size(); i++)
            System.out.print(arrli.get(i) + " ");
    }

}