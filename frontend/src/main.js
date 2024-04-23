window.onload = () => {
  console.log('loaded')
}

const randomWait = () => new Promise(res => setTimeout(res, Math.random() * 2000))

const imageContainer = document.getElementById('img-container')

let count = 0//
const addImage = (imageFile) => {
  console.log(imageFile)
  const img = document.createElement('img')
  img.height = 300
  img.src = URL.createObjectURL(imageFile)
  imageContainer.appendChild(img)

  count++//
  if (count > 30) img.style.visibility = 'hidden'//
}

window.getImageFilePaths = async () => {
  const imageFilePaths = await go.main.App.GetImageFilePaths()

  console.log(imageFilePaths)

  Promise.all(imageFilePaths.map(async path => {
    const response = await fetch(path)
    const imageFile = await response.blob()
    addImage(imageFile)
  }))
}
