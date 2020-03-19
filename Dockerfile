FROM busybox

COPY . /app
CMD /app/go-todo-list-app
