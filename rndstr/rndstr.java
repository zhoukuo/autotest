//javac Randomc.java
//java Randomc
import java.util.Random;

public class Randomc {

	public String RandomCharacter(int count) {
		Random random = new Random();
		StringBuilder str = new StringBuilder();

		for(int i=0; i<count; i++) {
			if (random.nextInt(2) == 0) {
				str = str.append((char)(random.nextInt(20902)+0x4e00));
			} else {
				str = str.append((char)(random.nextInt(94)+33));
			}
		}
		return str.toString();
	}

	public static void main(String[] args) {
		Randomc rc = new Randomc();
		long startTime = System.currentTimeMillis();
		rc.RandomCharacter(2000000);
		long endTime = System.currentTimeMillis();
		// Elapsed Time:73ms
		System.out.println("Elapsed Time:"+(endTime-startTime)+"ms");
	}
}