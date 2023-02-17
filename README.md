
# Forum

# Description
This project consists in creating a web forum that allows :
    communication between users.
    associating categories to posts.
    liking and disliking posts and comments.
    filtering posts.
SQLite
In order to store the data in the forum (like users, posts, comments, etc.) there is a need to use the database library SQLite. SQLite is a popular choice as an embedded database software for local/client storage in application software such as web browsers. It enables you to create a database as well as controlling it by using queries.

Authentication
In this segment the client must be able to register as a new user on the forum, by inputting their credentials. There is a need to create a login session to access the forum and be able to add posts and comments.Also cookies should be used to allow each user to have only one opened session. Each of this sessions must contain an expiration date. It is up to you to decide how long the cookie stays "alive". The use of UUID is a Bonus task.

Communication
In order for users to communicate between each other, they will have to be able to create posts and comments.

    Only registered users will be able to create posts and comments.
    When registered users are creating a post they can associate one or more categories to it.
        The implementation and choice of the categories is up to you.
    The posts and comments should be visible to all users (registered or not).
    Non-registered users will only be able to see posts and comments.

Likes and Dislikes
Only registered users will be able to like or dislike posts and comments. The number of likes and dislikes should be visible by all users (registered or not).

Filter
There is also a need to implement a filter mechanism, that will allow users to filter the displayed posts by :
    categories
    created posts
    liked posts

Finally, Docker, a containerization technology, must be used to make the project easier to launch.

# Run Locally

    1. Clone the project by running the following command in the terminal "git@git.01.alem.school:aliya.science/forum.git"
    2. Run the following command: "make run" and click on the URL address to go to the web page
    3. Then, the home page will be opened in which you can sign up to create your own account.
    4. After authorization, you need to sign in to go to your profile.
    5. Finally, it is up to you whether to post something, comment other's posts, like or dislike.
    6. You can also check for methods or follow the audit list.
 

## Authors

- [@aliya.science]
