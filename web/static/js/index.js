const modal = document.querySelector("[add-modal]");
const openModalButton = document.querySelector("[open-add-modal]");
const closeModalButton = document.querySelector("[close-add-modal]");

const submitBtn = document.querySelector("#add-task-btn")
const addTaskForm = document.querySelector("#add-task-form")

const description = document.querySelector("#description")
const title = document.querySelector("#title")

function toggleTask(id){
    axios.put(`/tasks/${id}`)
    .then(response => {
      console.log(response.data);
      // Optionally, you can refresh the page or update the UI to reflect the changes
      location.reload();
    })
    .catch(error => {
      console.error('There was an error toggling the task!', error);
    });
}

function deleteTask(id){
    axios.delete(`/tasks/${id}`)
    .then(response => {
      console.log(response.data);   
      location.reload();
    })
    .catch(error => {
      console.error('There was an error deleting the task!', error);
    });
}

openModalButton.addEventListener("click", () => {
  modal.showModal()
})

closeModalButton.addEventListener("click", () => {
  modal.close()
})

submitBtn.addEventListener("click", (event) => {
  const descriptionValue = document.querySelector("#description").value;
  const titleValue = document.querySelector("#title").value;

  axios({
    method: "post",
    url: "/tasks",
    data:{
      "title": titleValue,
      "description": descriptionValue,
    }
  })
  .then(response => {
    location.reload();
  })
  .catch(error => {
    console.error('There was an error adding the task!', error);
  });
})

const descriptionCharacterCount = document.querySelector("#description-character-count")
const titleCharacterCount = document.querySelector("#title-character-count")


title.addEventListener("keyup", () => {
  titleCharacterCount.innerText = title.value.length;
});

description.addEventListener("keyup", () => {
  descriptionCharacterCount.innerText = description.value.length
})