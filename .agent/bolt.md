## 2024-05-23 - [Setup]
**Vấn đề:** Chưa có file theo dõi hiệu năng và clean code.
**Giải pháp:** Khởi tạo file .agent/bolt.md để ghi lại các thay đổi.

## 2024-05-23 - [Memory & Bug Fix]
**Vấn đề:**
1. `GORMUserRepository.Update` sử dụng `map[string]interface{}` gây cấp phát bộ nhớ heap không cần thiết (Memory Allocation).
2. `GORMUserRepository.FindByID` không map các trường `FirstName`, `LastName`, `PhoneNumber`, `CountryCode`. Điều này dẫn đến việc `UpdateProfile` (gọi FindByID -> Update) vô tình xóa dữ liệu của các trường này (Critical Logic Bug).
3. `Update` ghi đè `created_at`, sai về mặt ngữ nghĩa (Immutable field).

**Giải pháp:**
1. Refactor `Update` để sử dụng struct `db.User` thay vì map. Giúp giảm allocation và tăng type safety.
2. Sử dụng `Omit("created_at")` để bảo vệ trường ngày tạo.
3. Bổ sung việc mapping đầy đủ các trường trong `FindByID`.
