const modal = document.querySelector("[add-modal]");
const openModalButton = document.querySelector("[open-add-modal]");
const closeModalButton = document.querySelector("[close-add-modal]");

const submitBtn = document.querySelector("#add-task-btn")
const addTaskForm = document.querySelector("#add-task-form")

function toggleTask(id){
    axios.put(`/api/tasks/${id}`)
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
    axios.delete(`/api/tasks/${id}`)
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
  event.preventDefault();
  console.log("clicked")
  const title = document.querySelector("#title").value;
  const description = document.querySelector('#description').value;

  axios({
    method: "post",
    url: "/api/tasks",
    data:{
      "title": title,
      "description": description,
    }
  })
  .then(response => {
    location.reload();
  })
  .catch(error => {
    console.error('There was an error adding the task!', error);
  });
})