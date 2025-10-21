package utils.java;

public class Pair {
        public final int row;
        public final int column;

        public Pair(int row, int column) {
                this.row = row;
                this.column = column;
        }

        // this is to control how the equality between objects
        // INFO: how to compare java objects?
        @Override
        public boolean equals(Object object) {
                if (this == object)
                        return true;

                if (!(object instanceof Pair))
                        return false;

                Pair pair = (Pair) object;
                return pair.row == row && pair.column == column;
        }

        @Override
        public String toString() {
                return "(" + row + ", " + column + ")";
        }
}
