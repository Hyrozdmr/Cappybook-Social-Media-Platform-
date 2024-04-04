// import { useState } from "react";

// const Comment = (props) => {
//   const [comment, setComment] = useState("");
//   const [comments, setComments] = useState(props.post.comments || []);

//   const handleCommentChange = (event) => {
//     setComment(event.target.value);
//   };

//   const handleSubmitComment = (event) => {
//     event.preventDefault();
//     // You can handle submitting the comment here
//     console.log("Submitting comment:", comment);
//     const newComment = { text: comment, id: Math.random().toString() }; // You can replace id generation with a better approach
//     setComments([...comments, newComment]); // Add the new comment to the list of comments
//     setComment(""); // Clear the comment input field after submission
//   };

//   return (
//     <div>
//       <h4>Comments:</h4>
//       <ul>
//         {comments.map((comment) => (
//           <li key={comment.id}>{comment.text}</li>
//         ))}
//       </ul>
//       <form onSubmit={handleSubmitComment}>
//         <input
//           type="text"
//           value={comment}
//           onChange={handleCommentChange}
//           placeholder="Add a comment..."
//         />
//         <button type="submit">Submit</button>
//       </form>
//     </div>
//   );
// };

// export default Comment;

import image from "/src/static/img/x-button.png";
import userImage from "/src/static/img/user_image.png";

const Comment= ({comment, post, onDelete, onSubmit, onChange}) => {

    const handleDeleteCommentClick = () => {
        onDelete(comment._id);
    };

    return (
        <div>
            <form onSubmit={onSubmit}>
                <div className="create-comment">
                    <input type="text" value={comment} onChange={onChange} />
                    <input role="submit-button" id="submit" type="submit" value="Submit" />
                </div>
            </form>
                <div className="comment-info">
                    <div className="comment-user">
                        <img className="user-image" src={userImage} alt="image" />
                        <p>{post.User.username}</p>
                    </div>
                    <p>{comment.message}</p>
                    <img className="delete-button" src={image} onClick={handleDeleteCommentClick} />
                </div>
        </div>
    );
};

export default Comment;





// const Comment = ({ post, comment, comments, onSubmit, onChange, onDelete }) => {
//     const handleCommentChange = (event) => {
//       onChange(event);
//     };
  
//     const handleSubmitComment = (event) => {
//       onSubmit(event);
//     };
  
//     const handleDeleteComment = (commentId) => {
//       onDelete(post._id, commentId);
//     };
  
//     return (
//       <div className="comment-container">
//         <h4>Comments:</h4>
//         <form onSubmit={handleSubmitComment}>
//           <input
//             type="text"
//             value={comment}
//             onChange={handleCommentChange}
//             placeholder="Add a comment..."
//           />
//           <button type="submit">Submit</button>
//         </form>
//         <ul>
//           {comments.map((comment) => (
//             <li key={comment.id}>
//               <div className="comment-info">
//                 <div className="comment-user">
//                   <p>{comment.user.username}</p>
//                 </div>
//                 <p>{comment.message}</p>
//                 <button onClick={() => handleDeleteComment(comment.id)}>Delete</button>
//               </div>
//             </li>
//           ))}
//         </ul>
       
//       </div>
//     );
//   };
  
//   export default Comment;