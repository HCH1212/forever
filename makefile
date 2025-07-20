.PHONY: linux
linux:
	fyne package -os linux -icon leimu.png

.PHONY: android
android:
	fyne package -os android -app-id online.hch1212.forever -icon leimu.png

.PHONY: ios
ios:
	fyne package -os ios -app-id online.hch1212.forever -icon leimu.png

.PHONY: install
install:
	fyne install -icon leimu.png
