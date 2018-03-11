FROM scratch
EXPOSE 8080
ADD main /
CMD /main -customer="${MY_CUSTOMER_NAME}" -pod="${MY_POD_NAME}"