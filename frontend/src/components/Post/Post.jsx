import React from 'react';
import "./Post.css";
import image from "/src/static/img/x-button.png";
import userImage from "/src/static/img/user_image.png";

const Post = ({ post, onLike, user, onDelete }) => {
    console.log("Rendering Post component...");

    const handleLikeClick = () => {
        onLike(post._id);
    };

    const handleDeleteClick = async () => {
        try {
            await onDelete(post._id);
        } catch (error) {
            console.error("Error deleting post:", error);
        }
    };

    return (
        <div className="post-container">
            <div className="post-info">
                <div className="post-user">
                    <img className="user-image" src={userImage} alt="image" />
                    <p>{user}</p>
                </div>
                <p>{post.message}</p>
                <p>Likes: {post.likes}</p>
                <button onClick={handleLikeClick}>Like</button>
            </div>
            <div className="post-image">
                <img className="delete-button" src={image} onClick={handleDeleteClick} alt=""/>
            </div>
        </div>
    );
};

export default Post;
