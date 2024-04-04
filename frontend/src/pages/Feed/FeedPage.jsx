import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { getPosts, createPosts, updatePostLikes, deletePosts} from "../../services/posts";
import Post from "../../components/Post/Post";
import Comment from "../../components/Comment/Comment";
import { getComments, createComments, deleteComments } from "../../services/comments";
import "./FeedPage.scss"

export const FeedPage = () => {
    const [posts, setPosts] = useState([]);
    const [post, setPost] = useState("");
    const [comments, setComments] = useState([]);
    const [comment, setComment] = useState("");
    const navigate = useNavigate();

    useEffect(() => {
        const token = localStorage.getItem("token");
        if (token) {
            getPosts(token)
                .then((data) => {
                    const sortedPosts = data.posts.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
                    setPosts(sortedPosts);
                    localStorage.setItem("token", data.token);
                })
                .catch((err) => {
                    console.error(err);
                    navigate("/login");
                });
            getComments(token)
                .then((data) => {
                    const sortedComments = data.comments.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
                    setComments(sortedComments);
                    localStorage.setItem("token", data.token);
                })
                .catch((err) => {
                    console.error(err);
                });
    }
    }, [navigate]);

    const handleLike = async (postId) => {
        try {
            await updatePostLikes(token, postId);
            const updatedPosts = await getPosts(token);
            const sortedPosts = updatedPosts.posts.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
            setPosts(sortedPosts);
        } catch (err) {
            console.error(err);
        }
    };
      const token = localStorage.getItem("token");
    if (!token) {
      navigate("/login");
      return;
    }

    const handleDelete = async (postId) => {
      try {
          await deletePosts(token, postId);
          const updatedPosts = await getPosts(token);
          const sortedPosts = updatedPosts.posts.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
          setPosts(sortedPosts);
      } catch (err) {
          console.error(err);
      }
  };
  

    const handleSubmitPost = async (event) => {
      event.preventDefault();
      try {
        await createPosts(token, post);
        const updatedPosts = await getPosts(token);
        const sortedPosts = updatedPosts.posts.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
        setPosts(sortedPosts);
        setPost("");
        localStorage.setItem("token", updatedPosts.token);
      } catch (err) {
        console.error(err);
      }
    };

    const handleSubmitComment = async (event, postId) => {
      event.preventDefault();
        try {
            console.log("Token:", token);
            console.log("Comment:", comment);
            console.log("Post ID:", postId);
            await createComments(token, comment, postId);
            const updatedComments = await getComments(postId, token);
            const sortedComments = updatedComments.comments.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
            setComments(sortedComments);
            setComment("");
            localStorage.setItem("token", updatedComments.token);
        } catch (err) {
            console.error(err);
        }
    };

    const handleDeleteComment = async (postId, commentId) => {
      try {
          await deleteComments(token, postId, commentId);
          const updatedComments = await getComments(token);
            const sortedComments = updatedComments.comments.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
            setComments(sortedComments);
      } catch (err) {
          console.error(err);
      }
  };

    const handleCommentChange = (event) => {
      setComment(event.target.value);
  };

    const handlePostChange = (event) => {
      setPost(event.target.value);
    };

    return (
      <div className="feed-container">
        <h2>Posts</h2>
        <form onSubmit={handleSubmitPost}>
          <div className="create-post">
            <input type="text" value={post} onChange={handlePostChange} />
            <input role="submit-button" id="submit" type="submit" value="Submit" />
          </div>
        </form>
        <div className="feed-all-posts" role="feed">
          {posts.map((post) => (
            <div className="feed-post" key={post._id}>
              <Post post={post} onDelete={handleDelete} onLike={handleLike} user={post.User.username} />
  
              <form onSubmit={(e) => handleSubmitComment(e, post._id)}>
                <div className="create-comment">
                  <input type="text" value={comment} onChange={handleCommentChange} />
                  <input role="submit-button" id="submit" type="submit" value="Submit" />
                </div>
              </form>
  
              {comments
                .filter((comment) => comment.postId === post._id)
                .map((comment) => (
                  <div className="feed-comment" key={comment._id}>
                    <Comment comment={comment} onDelete={() => handleDeleteComment(post._id, comment._id)} />
                  </div>
                ))}
            </div>
          ))}
        </div>
      </div>
    );
  };

