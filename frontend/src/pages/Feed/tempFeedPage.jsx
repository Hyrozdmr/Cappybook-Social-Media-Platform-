import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";

import { getPosts } from "../../services/posts";
import { createPosts } from "../../services/posts";
import Post from "../../components/Post/Post";
import Comment from "../../components/Comment/Comment";

export const FeedPage = () => {
    const [posts, setPosts] = useState([]);
    const [post, setPost] = useState("");
    const navigate = useNavigate();

    useEffect(() => {
        const token = localStorage.getItem("token");
        if (token) {
            getPosts(token)
                .then((data) => {
                    const sortedPosts = data.posts.sort(
                        (a, b) => new Date(b.created_at) - new Date(a.created_at)
                    );
                    setPosts(sortedPosts);
                })
                .catch((err) => {
                    console.error(err);
                    navigate("/login");
                });
        }
    }, [navigate]);

    const handleLike = async (postId) => {
        try {
            // Send request to backend to update like count for the post with postId
            console.log(`Like button clicked for post with ID: ${postId}`);
        } catch (err) {
            console.error(err);
        }
    };

    const handleSubmit = async (event) => {
        event.preventDefault();
        try {
            const createdPostResponse = await createPosts(token, post);
            const updatedPosts = await getPosts(token);
            const sortedPosts = updatedPosts.posts.sort(
                (a, b) => new Date(b.created_at) - new Date(a.created_at)
            );
            setPosts(sortedPosts);
            setPost("");
            localStorage.setItem("token", updatedPosts.token);
        } catch (err) {
            console.error(err);
        }
    };

    const handlePostChange = (event) => {
        setPost(event.target.value);
    };

    return (
        <div className="container">
            <h2>Posts</h2>
            <div className="feed" role="feed">
                {posts.map((post) => (
                    <div key={post._id}>
                        <Post post={post} />
                        {/* Like button */}
                        <button onClick={() => handleLike(post._id)}>Like</button>
                        {/* Render the Comment component for each post */}
                        <Comment post={post} />
                    </div>
                ))}
            </div>
            <form onSubmit={handleSubmit}>
                <div className="create-post">
                    <input type="text" value={post} onChange={handlePostChange} />
                    <input role="submit-button" id="submit" type="submit" value="Submit" />
                </div>
            </form>
        </div>
    );
};
