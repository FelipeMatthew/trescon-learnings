import cv2
import numpy as np

video = cv2.VideoCapture('escalator.mp4')
counter = 0
opened = True

while True:
    ret, img = video.read()  # Changed from video.open() to video.read()
    if not ret:  # Check if frame is not read correctly
        break
    
    img = cv2.resize(img, (1100, 720))
    imgGray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)  # Changed COLOR_RGB2GRAY to COLOR_BGR2GRAY

    # Th = Threshold - Limite
    imgTh = cv2.adaptiveThreshold(imgGray, 255, cv2.ADAPTIVE_THRESH_GAUSSIAN_C, cv2.THRESH_BINARY_INV, 11, 12)
    kernel = np.ones((8, 8), np.uint8)
    # Imagem dilatada
    imgDil = cv2.dilate(imgTh, kernel, iterations=2)  # Changed interations to iterations

    # Rectangle pixels
    x, y, w, h = 490, 320, 30, 150

    recortes = imgDil[y:y + h, x:x + w]
    brancos = cv2.countNonZero(recortes)

    if brancos > 4000 and opened == True:
        counter += 1
    if brancos < 4000:
        opened = True
    else:
        opened = False

    if opened == False:
        cv2.rectangle(img, (x, y), (x + w, y + h), (0, 255, 0), 4)
    else:
        cv2.rectangle(img, (x, y), (x + w, y + h), (255, 0, 255), 4)

    cv2.rectangle(imgTh, (x, y), (x + w, y + h), (255, 255, 255), 6)

    # Colocando count non value
    cv2.putText(img, str(brancos), (x - 30, y - 50), cv2.FONT_HERSHEY_SIMPLEX, 1, (255, 255, 255), 1)
    cv2.rectangle(img, (577, 155), (577 + 88, 155 + 85), (255, 255, 255), -1)
    cv2.putText(img, str(counter), (x + 100, y), cv2.FONT_HERSHEY_SIMPLEX, 3, (255, 0, 0), 5)

    cv2.imshow('Original video', img)
    cv2.imshow('Video Th', cv2.resize(imgTh, (600, 500)))

    if cv2.waitKey(20) & 0xFF == ord('q'):  # Exit loop if 'q' is pressed
        break

video.release()
cv2.destroyAllWindows()
