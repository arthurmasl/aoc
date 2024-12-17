const text = await Bun.file(`${import.meta.dir}/example.txt`).text();
const lines = text.trim().split('\n');

function sumOfCalibrationValues(calibrationDocument) {
  let sum = 0;

  // Loop through each line in the calibration document array
  for (let i = 0; i < calibrationDocument.length; i++) {
    let line = calibrationDocument[i];

    // Extracting the first and last digits if the line has at least two characters
    if (line.length >= 2) {
      let firstDigit = parseInt(line.charAt(0));
      let lastDigit = parseInt(line.charAt(line.length - 1));

      // Ensure both extracted characters are digits
      if (!isNaN(firstDigit) && !isNaN(lastDigit)) {
        let calibrationValue = firstDigit * 10 + lastDigit;
        sum += calibrationValue;
      }
    }
  }

  return sum;
}

console.log(sumOfCalibrationValues(lines));
