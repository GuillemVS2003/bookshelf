package vidal.bookshelf.repositories;

import org.springframework.data.jpa.repository.JpaRepository;
import vidal.bookshelf.models.Book;

public interface BookRepository extends JpaRepository<Book, Long> {

}
