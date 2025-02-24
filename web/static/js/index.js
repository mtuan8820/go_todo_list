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

const openModalButton = document.querySelector("[open-add-modal]");
const closeModalButton = document.querySelector("[close-add-modal]");
const modal = document.querySelector("[add-modal]");

openModalButton.addEventListener("click", () => {
  modal.showModal()
})

closeModalButton.addEventListener("click", () => {
  modal.close()
})