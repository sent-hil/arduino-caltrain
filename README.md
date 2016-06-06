# arduino-caltrain

Script to get caltrain timings and send it to Arduino using serial. This script
has only been tested to work on Arduino UNO and assumes the device is at
`/dev/tty.usbmodem1421`.

You'll also need to upload `serialdisplay.ino` to the Arduino device first.

## Crontab

Recommended to use the following crontab to run the script every weekday from
4pm to 9pm on 1 minute interval.

    * 0,16,17,18,19,20,21 * * 1-5 /path/to/arduino-caltrain > /var/log/arduino-caltrain.log
