import cv2

# Load the pre-trained face detection model
face_cascade = cv2.CascadeClassifier(cv2.data.haarcascades + 'haarcascade_frontalface_default.xml')

# Load the pre-trained age, emotion, and gender detection models
age_net = cv2.dnn.readNetFromCaffe('deploy_age.prototxt', 'age_net.caffemodel')
emotion_net = cv2.dnn.readNetFromCaffe('deploy_emotion.prototxt', 'emotion_net.caffemodel')
gender_net = cv2.dnn.readNetFromCaffe('deploy_gender.prototxt', 'gender_net.caffemodel')

# Open the default camera
cap = cv2.VideoCapture(0)

while True:
    # Read the current frame from the camera
    ret, frame = cap.read()

    # Convert the frame to grayscale
    gray = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)

    # Detect faces in the frame
    faces = face_cascade.detectMultiScale(gray, scaleFactor=1.1, minNeighbors=5, minSize=(30, 30))

    # Process each detected face
    for (x, y, w, h) in faces:
        # Extract the face region of interest
        face_roi = frame[y:y+h, x:x+w]

        # Resize the face ROI to match the input size of the age, emotion, and gender detection models
        face_blob = cv2.dnn.blobFromImage(face_roi, 1.0, (227, 227), (78.4263377603, 87.7689143744, 114.895847746), swapRB=False)

        # Pass the face blob through the age, emotion, and gender detection models
        age_net.setInput(face_blob)
        age_preds = age_net.forward()
        emotion_net.setInput(face_blob)
        emotion_preds = emotion_net.forward()
        gender_net.setInput(face_blob)
        gender_preds = gender_net.forward()

        # Get the predicted age, emotion, and gender labels
        age_label = int(age_preds[0].argmax())
        emotion_label = int(emotion_preds[0].argmax())
        gender_label = int(gender_preds[0].argmax())

        # Draw rectangles around the detected faces
        cv2.rectangle(frame, (x, y), (x+w, y+h), (0, 255, 0), 2)

        # Display the age, emotion, and gender labels
        label = f"Age: {age_label}, Emotion: {emotion_label}, Gender: {gender_label}"
        cv2.putText(frame, label, (x, y-10), cv2.FONT_HERSHEY_SIMPLEX, 0.9, (0, 255, 0), 2)

    # Display the frame with detected faces and labels
    cv2.imshow('Face Recognition', frame)

    # Break the loop if 'q' is pressed
    if cv2.waitKey(1) & 0xFF == ord('q'):
        break

# Release the camera and destroy all windows
cap.release()
cv2.destroyAllWindows()
