import {
  createTodo,
  deleteAllTodos,
  deleteTodo,
  fetchTaskById,
  handleEdit,
  updateAllTodos,
  updateStatus,
} from "./services.js";
import { splitId } from "./utils.js";

// CREATE TODO
let createBtn = document.querySelector(".todo__logo");
createBtn.addEventListener("click", async () => {
  const inputTask = document.getElementById("todo").value;

  if (inputTask === "") {
    alert("Task should not be emtpy!");
    return;
  }

  const payload = {
    task: inputTask,
    status: "PENDING",
  };

  await createTodo(payload);
});

// COMPLETE ALL TODOS
let completeAllBtn = document.querySelector(".todo__complete");
completeAllBtn.addEventListener("click", async () => {
  await updateAllTodos();
});

// DELETE ALL TODOS
let deleteAllBtn = document.querySelector(".todo__delete");
deleteAllBtn.addEventListener("click", async () => {
  await deleteAllTodos();
});

// EDIT ICON CLICK FOR POPUP
let editBtns = document.querySelectorAll(".fa-solid.fa-pen-to-square");
editBtns.forEach((editBtn) => {
  editBtn.addEventListener("click", async (event) => {
    const id = event.target.id;
    const taskInfo = splitId(id);

    const data = await fetchTaskById(taskInfo[1]);

    document.querySelector(".popup__container").style.display = "block";

    const statusRadios = document.querySelectorAll('input[name="status"]');
    statusRadios.forEach((radio) => {
      if (radio.value === data.payload.status) {
        radio.checked = true;
      }
    });

    const editTask = document.getElementById("editTask");
    editTask.value = data.payload.task;

    const form = document.querySelector("form");
    form.setAttribute("data-task-id", data.taskId);
  });
});

// UPDATE STATUS
let statusBtns = document.querySelectorAll(".todo__list-task-check");
statusBtns.forEach((statusBtn) => {
  statusBtn.addEventListener("click", async (event) => {
    const id = event.target.id;
    const taskInfo = splitId(id);

    const taskName = event.currentTarget
      .closest(".todo__list-task")
      .querySelector(".todo__list-task-name")
      .textContent.trim();

    let payload = {
      task: taskName,
      status: taskInfo[0] === "PENDING" ? "COMPLETED" : "PENDING",
    };

    await updateStatus(taskInfo[1], payload);
  });
});

// DELETE TODO
let deleteBtns = document.querySelectorAll(".fa-solid.fa-trash");
deleteBtns.forEach((deleteBtn) => {
  deleteBtn.addEventListener("click", async (event) => {
    const id = event.target.id;
    const taskInfo = splitId(id);

    await deleteTodo(taskInfo[1]);
  });
});

function closePopup() {
  document.querySelector(".popup__container").style.display = "none";
}

// MAKING IT GLOBAL TO ACCESS DIRECTLY FROM HTML
window.closePopup = closePopup;
window.handleEdit = handleEdit;
