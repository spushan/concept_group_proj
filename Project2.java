import java.util.Scanner;

class Project2 {
    
    public static class Matrix extends Thread {
        
        private int row;
        private int col;
        private double p;
        private double[][] matrix;

        public Matrix(int row, int col, double p){
            
            this.row = row;
            this.col = col;
            this.p = p;
            matrix = new double[row][col];
            
            for(int i=0; i<row; i++){
                for(int j=0; j<col; j++){
                    matrix[i][j] = p;
                }
            }
        }

        public int getRow(){
            return row;
        }

        public int getCol(){
            return col;
        }

        public double getValue(int row, int col){
            return matrix[row][col];
        }

        public void print(){
            System.out.println("");
            for(int i=0; i<row; i++){
                for(int j=0; j<col; j++){
                    System.out.print(matrix[i][j] + " ");
                }
                System.out.println("");
            }
        }
        
        
        public Matrix mul(Matrix other){

            Matrix result = new Matrix(row, other.getCol(), 0);

            for(int i=0; i<row; i++){
                for(int j=0; j<other.getRow(); j++){
                    double ans = 0;
                    for(int k=0; k<col;k++){
                        ans += matrix[i][k] * other.getValue(k, j); 
                    }
                    result.matrix[i][j] = ans;
                }
            }
            return result;
        }
        
        public Matrix pmul(Matrix other){
            
            Matrix result = new Matrix(row, other.getCol(), 0);
            
            for(int i=0; i<row; i++){
                final int s = i;
                new Thread (()->{
                    for(int j=0; j<other.getRow(); j++){
                        double ans = 0;
                        for(int k=0; k<col;k++){
                            ans += matrix[s][k] * other.getValue(k, j); 
                        }
                        result.matrix[s][j] = ans;
                    }
                }).start();
            }
            return result;
        }
    }

    public static  Matrix[] maker(int row, int col, int row2, int col2, double p, Scanner scanner) {
        
        
        Matrix matArr[] = new Matrix[3];
        
        if(row==0 || col==0 || row2==0 || col2==0){
            throw new ArithmeticException("0 for size not allowed");
        }
        if(col != row2){
            throw new ArithmeticException("Column of Matrix 1 and Row of Matrix 2 must be Equal");
        }
        
        Matrix mat1 = new Matrix(row, col, p);
        Matrix mat2 = new Matrix(row2, col2, p);
        Matrix answ = new Matrix(row, col2, 0);
        matArr[0] = mat1;
        matArr[1] = mat2;
        matArr[2] = answ;
        return matArr;
    }
    public static void main(String[] args) {
        
        int row, col, row2,col2;
        double p;
        String inp;
        Scanner scanner = new Scanner(System.in);
        Matrix matArr[], mat1, mat2, answ;

        
        while(true) {
            System.out.print("\nEnter First Matrix Row: ");
            row = scanner.nextInt();
            System.out.print("\nEnter First Matrix Column: ");
            col = scanner.nextInt();
            System.out.print("\nEnter Second Matrix Row: ");
            row2 = scanner.nextInt();
            System.out.print("\nEnter Second Matrix Column: ");
            col2 = scanner.nextInt();
            System.out.print("\nEnter Number to populate with: ");
            p = scanner.nextDouble();

            matArr = maker(row, col, row2, col2, p, scanner);
            mat1 = matArr[0];
            mat2 = matArr[1]; 
            answ = matArr[2];
            System.out.println("\nNormal Multiply");
            long start = System.currentTimeMillis();
            answ = mat1.mul(mat2);
            long end =System.currentTimeMillis();
            System.out.println("Elapse Time " + (end - start) + " milliseconds");
            if (row < 10 && col < 10) {
                answ.print();
            }

            System.out.println("\nMultithread Multiply");
            start = System.currentTimeMillis();
            answ = mat1.mul(mat2);
            end = System.currentTimeMillis();
            System.out.println("Elapse Time " + (end - start) + " milliseconds");
            if (row < 10 && col < 10) {
                answ.print();
            }
            System.out.println("Enter New Matrix? Y/n: ");
            inp = scanner.next();
            
            if(inp == "n") {
                System.exit(1);
                scanner.close();
            }



        }
        



        
        

    }
}
